import { useState, useEffect } from 'react'
import { Card, CardTitle, CardBody, Spinner } from '@patternfly/react-core'
import { performHandshake } from './handshake'

function App() {
  const [handshakeStatus, setHandshakeStatus] = useState('idle')
  const [pairingStatus, setPairingStatus] = useState('idle')
  const [pairingRequestId, setPairingRequestId] = useState(null)
  const [pairingErrorMessage, setPairingErrorMessage] = useState(null)
  const [approvalStatus, setApprovalStatus] = useState('idle')


  useEffect(() => {
    const runHandshake = async () => {
      setHandshakeStatus('loading')
      try {
        // Extract token from URL fragment (e.g., #token=value or #value)
        const fragment = window.location.hash
        const token = fragment ? fragment.replace(/^#\/?(?:token=)?/, '') : undefined
        if (token) {
          console.log('Using auth token from URL fragment')
        }

        // Construct WebSocket URL from current host
        const wsProtocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
        const gatewayUrl = import.meta.env.VITE_GATEWAY_URL || `${wsProtocol}//${window.location.host}`
        await performHandshake({
          clientId: 'openclaw-control-ui',
          clientVersion: '2026.5.3',
          mode: 'webchat',
          role: 'operator',
          scopes: [
            'operator.admin',
            'operator.approvals',
            'operator.pairing',
            'operator.read',
            'operator.write'
          ],
          platform: 'MacIntel',
          deviceFamily: 'webchat',
          token,
          gatewayUrl
        })
        setHandshakeStatus('success')
        setPairingStatus('success')
        setApprovalStatus('approved')
      } catch (err) {
        console.error('Handshake error:', err.code, err.message, err)

        // NOT_PAIRED is expected - device ID was generated, now needs pairing
        if (err.code === 'NOT_PAIRED') {
          setHandshakeStatus('success')
          const requestId = err.details?.requestId
          if (requestId) {
            console.log('Device not paired, pairing request ID:', requestId)
            setPairingRequestId(requestId)
          } else {
            console.error('NOT_PAIRED error missing requestId in details')
            setHandshakeStatus('error')
            setPairingErrorMessage('Incomplete error information from server')
          }
        } else {
          // Other errors are actual failures
          setHandshakeStatus('error')
          setPairingErrorMessage(err.message || 'Handshake failed')
        }
      }
    }
    runHandshake()
  }, [])

  // Submit pairing request when requestId is available
  useEffect(() => {
    const submitPairingRequest = async () => {
      if (!pairingRequestId) return

      setPairingStatus('pending')
      try {
        // Create abort controller for timeout
        const controller = new AbortController()
        const timeoutId = setTimeout(() => controller.abort(), 10000) // 10 second timeout

        const response = await fetch('pairing-requests', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ requestId: pairingRequestId }),
          signal: controller.signal
        })

        clearTimeout(timeoutId)

        if (!response.ok) {
          const errorData = await response.json().catch(() => ({}))
          throw new Error(errorData.error || `HTTP ${response.status}`)
        }

        console.log('Pairing request submitted successfully, starting status polling')
        setPairingStatus('progressing')
        setApprovalStatus('pending')
      } catch (err) {
        console.error('Pairing submission failed:', err)
        if (err.name === 'AbortError') {
          setPairingErrorMessage('Request timed out - please try again')
        } else {
          setPairingErrorMessage(err.message || 'Failed to submit pairing request')
        }
        setPairingStatus('error')
      }
    }

    submitPairingRequest()
  }, [pairingRequestId])

  // Poll pairing status after submission
  useEffect(() => {
    if (!pairingRequestId || pairingStatus !== 'progressing') return



    const startTime = Date.now()
    const POLL_INTERVAL = 1000 // 1 second
    const TIMEOUT_MS = 30000 // 30 seconds

    const checkStatus = async () => {
      try {
        const response = await fetch(`pairing-requests/${pairingRequestId}`, {
          method: 'GET'
        })

        if (response.status === 200) {
          // Pairing approved
          console.log('Pairing approved')
          setPairingStatus('success')
          setApprovalStatus('approved')

          return true
        } else if (response.status === 202) {
          // Still pending
          console.log('Pairing pending, continuing to poll')
          return false
        } else {
          // Error
          console.error('Unexpected status during polling:', response.status)
          setPairingStatus('error')
          setApprovalStatus('error')
          setPairingErrorMessage('Failed to check pairing status')

          return true
        }
      } catch (err) {
        console.error('Error polling pairing status:', err)
        setPairingStatus('error')
        setApprovalStatus('error')
        setPairingErrorMessage('Network error while checking status')
        return true
      }
    }

    const pollInterval = setInterval(async () => {
      const elapsed = Date.now() - startTime

      if (elapsed >= TIMEOUT_MS) {
        console.warn('Polling timeout after 30 seconds')
        setPairingStatus('error')
        setApprovalStatus('timeout')
        setPairingErrorMessage('Pairing approval timed out')
        clearInterval(pollInterval)
        return
      }

      const shouldStop = await checkStatus()
      if (shouldStop) {
        clearInterval(pollInterval)
      }
    }, POLL_INTERVAL)

    // Cleanup on unmount
    return () => {
      clearInterval(pollInterval)
    }
  }, [pairingRequestId, pairingStatus])

  // Navigation handler
  const navigateToOpenClaw = () => {
    const fragment = window.location.hash
    const tokenMatch = fragment.match(/token=([^&]+)/)

    if (!tokenMatch) {
      console.warn('No token found in URL fragment during navigation')
    }

    const token = tokenMatch ? tokenMatch[1] : ''
    const protocol = window.location.protocol
    const host = window.location.host
    const newUrl = token ? `${protocol}//${host}#token=${token}` : `${protocol}//${host}`

    console.log('Navigating to OpenClaw:', newUrl)
    window.location.href = newUrl
  }

  // Auto-redirect when pairing is approved
  useEffect(() => {
    if (approvalStatus === 'approved') {
      navigateToOpenClaw()
    }
  }, [approvalStatus])

  const isError = handshakeStatus === 'error' || pairingStatus === 'error'

  const statusLabel = (() => {
    if (isError) return pairingErrorMessage || 'An error occurred'
    if (approvalStatus === 'approved') return 'Redirecting to OpenClaw...'
    if (pairingStatus === 'pending' || pairingStatus === 'progressing') return 'Pairing device with OpenClaw...'
    return 'Generating device id...'
  })()

  return (
    <Card>
      <CardTitle>Device Pairing</CardTitle>
      <CardBody style={{ textAlign: 'center' }}>
        {!isError && <Spinner aria-label="Loading" />}
        <p style={{ marginTop: '16px' }}>{statusLabel}</p>
      </CardBody>
    </Card>
  )
}

export default App
