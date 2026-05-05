import { useState, useEffect } from 'react'
import { Card, CardTitle, CardBody, ProgressStepper, ProgressStep } from '@patternfly/react-core'
import { performHandshake } from './handshake'

function App() {
  const [handshakeStatus, setHandshakeStatus] = useState('idle')
  const [deviceToken, setDeviceToken] = useState(null)

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
        const result = await performHandshake({
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
        setDeviceToken(result.deviceToken)
        setHandshakeStatus('success')
      } catch (err) {
        const errorMsg = err.message || 'Handshake failed'
        console.error('Handshake failed:', errorMsg, err)
        setHandshakeStatus('error')
      }
    }
    runHandshake()
  }, [])

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
          <ProgressStep id="step-2" titleId="step-2-title">
            Pair device with OpenClaw
          </ProgressStep>
        </ProgressStepper>
      </CardBody>
    </Card>
  )
}

export default App
