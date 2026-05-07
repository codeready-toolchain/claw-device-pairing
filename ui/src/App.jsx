import { useState, useEffect } from 'react'
import { Card, CardTitle, CardBody, ProgressStepper, ProgressStep, Button } from '@patternfly/react-core'
import { performHandshake } from './handshake'

function App() {
  const [handshakeStatus, setHandshakeStatus] = useState('idle')
  const [pairingStatus, setPairingStatus] = useState('idle')
  const [pairingRequestId, setPairingRequestId] = useState(null)
  const [pairingErrorMessage, setPairingErrorMessage] = useState(null)
  const [approvalStatus, setApprovalStatus] = useState('idle')
  const [isPolling, setIsPolling] = useState(false)

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

    setIsPolling(true)
    setApprovalStatus('pending')

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
          setIsPolling(false)
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
          setIsPolling(false)
          return true
        }
      } catch (err) {
        console.error('Error polling pairing status:', err)
        setPairingStatus('error')
        setApprovalStatus('error')
        setPairingErrorMessage('Network error while checking status')
        setIsPolling(false)
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
        setIsPolling(false)
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
      setIsPolling(false)
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

  return (
    <Card>
      <CardTitle>Device Pairing</CardTitle>
      <CardBody>
        <ProgressStepper isVertical={true}>
          <ProgressStep
            id="step-1"
            titleId="step-1-title"
            variant={handshakeStatus === 'success' ? 'success' : handshakeStatus === 'error' ? 'danger' : undefined}
          >
            Generate device id
          </ProgressStep>
          <ProgressStep
            id="step-2"
            titleId="step-2-title"
            variant={
              pairingStatus === 'success' ? 'success' :
              pairingStatus === 'error' ? 'danger' :
              pairingStatus === 'pending' || pairingStatus === 'progressing' ? 'info' :
              undefined
            }
            description={pairingStatus === 'error' && pairingErrorMessage ? pairingErrorMessage : undefined}
          >
            Pair device with OpenClaw
          </ProgressStep>
        </ProgressStepper>
        <Button
          onClick={navigateToOpenClaw}
          isDisabled={approvalStatus !== 'approved'}
          style={{ marginTop: '20px' }}
        >
          Go to OpenClaw
        </Button>
      </CardBody>
    </Card>
  )
}

export default App
