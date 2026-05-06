import { useState, useEffect } from 'react'
import { Card, CardTitle, CardBody, ProgressStepper, ProgressStep } from '@patternfly/react-core'
import { performHandshake } from './handshake'

function App() {
  const [handshakeStatus, setHandshakeStatus] = useState('idle')
  const [pairingStatus, setPairingStatus] = useState('idle')
  const [pairingRequestId, setPairingRequestId] = useState(null)
  const [pairingErrorMessage, setPairingErrorMessage] = useState(null)

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
          clientVersion: '1.0.0',
          mode: 'webchat',
          role: 'operator',
          scopes: [
            'operator.admin',
            'operator.read',
            'operator.write',
            'operator.approvals',
            'operator.pairing'
          ],
          platform: 'web',
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
        const response = await fetch('pairing-requests', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ requestId: pairingRequestId })
        })

        if (!response.ok) {
          const errorData = await response.json().catch(() => ({}))
          throw new Error(errorData.error || `HTTP ${response.status}`)
        }

        console.log('Pairing request submitted successfully')
        setPairingStatus('success')
      } catch (err) {
        console.error('Pairing submission failed:', err)
        setPairingErrorMessage(err.message || 'Failed to submit pairing request')
        setPairingStatus('error')
      }
    }

    submitPairingRequest()
  }, [pairingRequestId])

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
            variant={pairingStatus === 'success' ? 'success' : pairingStatus === 'error' ? 'danger' : pairingStatus === 'pending' ? 'info' : undefined}
            description={pairingStatus === 'error' && pairingErrorMessage ? pairingErrorMessage : undefined}
          >
            Pair device with OpenClaw
          </ProgressStep>
        </ProgressStepper>
      </CardBody>
    </Card>
  )
}

export default App
