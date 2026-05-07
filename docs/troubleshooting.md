# Troubleshooting Guide

## Common Startup Issues

### Issue: "failed to load in-cluster configuration"

**Symptoms:**
- Pod is in CrashLoopBackOff state
- Server logs show: "failed to initialize Kubernetes client"
- Error message includes: "unable to load in-cluster configuration, KUBERNETES_SERVICE_HOST and KUBERNETES_SERVICE_PORT must be defined"

**Causes:**
1. Server is not running in a Kubernetes cluster
2. In-cluster service account credentials are not mounted
3. Required environment variables are missing

**Solutions:**

1. Verify the pod is running inside a Kubernetes cluster:
   ```bash
   kubectl get pods -n <namespace>
   ```

2. Check service account is mounted:
   ```bash
   kubectl exec <pod-name> -n <namespace> -- ls -la /var/run/secrets/kubernetes.io/serviceaccount/
   ```
   Should show: `ca.crt`, `namespace`, `token`

3. Verify environment variables are set:
   ```bash
   kubectl exec <pod-name> -n <namespace> -- env | grep KUBERNETES
   ```
   Should show: `KUBERNETES_SERVICE_HOST` and `KUBERNETES_SERVICE_PORT`

4. Check the deployment has a service account configured:
   ```bash
   kubectl get deployment claw-device-pairing -n <namespace> -o yaml | grep serviceAccountName
   ```

**Note:** The server cannot run outside of a Kubernetes cluster. Local development requires deploying to a development cluster.

---

## Common CR Creation Issues

### Issue: "requestID cannot be sanitized to valid DNS-1123 label"

**Symptoms:**
- Pairing requests fail with HTTP 500
- Server logs show sanitization error

**Cause:**
The request ID contains only special characters that cannot be converted to a valid Kubernetes resource name.

**Solutions:**
- Ensure request IDs contain at least some alphanumeric characters
- Review the request ID generation logic in the client
- Valid characters: lowercase letters, numbers, hyphens (must start/end with alphanumeric)

**Examples:**
- ✅ Valid: `request-123`, `abc-def-456`, `my-request`
- ❌ Invalid: `@#$%`, `---`, `!!!`

---

### Issue: "failed to create pairing request CR: forbidden"

**Symptoms:**
- Pairing requests fail with HTTP 500
- Server logs: "forbidden: User cannot create resource"

**Cause:**
The service account lacks RBAC permissions to create `ClawDevicePairingRequest` resources.

**Solutions:**

1. Verify the service account exists:
   ```bash
   kubectl get serviceaccount claw-device-pairing -n <namespace>
   ```

2. Check RBAC permissions:
   ```bash
   kubectl get role claw-device-pairing -n <namespace> -o yaml
   kubectl get rolebinding claw-device-pairing -n <namespace> -o yaml
   ```

3. Apply the correct RBAC configuration (see `deployment.md`)

4. Restart the pods to pick up new RBAC settings:
   ```bash
   kubectl rollout restart deployment claw-device-pairing -n <namespace>
   ```

---

### Issue: Server fails to start - missing environment variables

**Symptoms:**
- Pod is in CrashLoopBackOff state
- Server logs show panic: "required environment variables not set: [NAMESPACE]" or "[CLAW_INSTANCE]" or both
- Pod restarts continuously

**Cause:**
The `NAMESPACE` and/or `CLAW_INSTANCE` environment variables are not set. Both are required and the server will panic at startup if either is missing.

**Solutions:**

1. Check pod environment variables:
   ```bash
   kubectl exec <pod-name> -n <namespace> -- env | grep -E "NAMESPACE|CLAW_INSTANCE"
   ```

2. Verify deployment manifest has correct Downward API configuration (see `deployment.md`):
   ```yaml
   env:
   - name: NAMESPACE
     valueFrom:
       fieldRef:
         fieldPath: metadata.namespace
   - name: CLAW_INSTANCE
     valueFrom:
       fieldRef:
         fieldPath: metadata.labels['claw.sandbox.redhat.com/instance']
   ```

3. Ensure pod has the `claw.sandbox.redhat.com/instance` label:
   ```bash
   kubectl get pod <pod-name> -n <namespace> --show-labels
   ```

4. If label is missing, update the deployment:
   ```yaml
   template:
     metadata:
       labels:
         claw.sandbox.redhat.com/instance: "your-instance-name"
   ```

5. After fixing the configuration, redeploy:
   ```bash
   kubectl apply -f deploy/kubernetes/deployment.yaml
   ```

---

### Issue: CR created but operator doesn't process it

**Symptoms:**
- CR exists but pairing status remains "pending"
- Operator logs don't show processing activity

**Possible Causes:**

1. **Selector mismatch**: CR selector doesn't match operator's instance label
   ```bash
   # Check CR selector
   kubectl get clawdevicepairingrequest <cr-name> -n <namespace> -o yaml | grep -A5 selector
   
   # Check operator pod labels
   kubectl get pod -l app=claw-operator -n <namespace> --show-labels
   ```

2. **Operator not watching**: Operator might not be configured to watch this namespace
   - Check operator logs for namespace watch configuration
   - Verify operator has RBAC to list/watch CRs in this namespace

3. **CR in different namespace**: Pairing server and operator in different namespaces
   - Verify both are in the same namespace or operator watches multiple namespaces

---

## Diagnostic Commands

### Check CR creation
```bash
# List all pairing request CRs
kubectl get clawdevicepairingrequests -n <namespace>

# Get details of specific CR
kubectl get clawdevicepairingrequest <cr-name> -n <namespace> -o yaml

# Watch CRs being created in real-time
kubectl get clawdevicepairingrequests -n <namespace> --watch
```

### Check server logs
```bash
# Follow server logs
kubectl logs -f deployment/claw-device-pairing -n <namespace>

# Search for CR creation logs
kubectl logs deployment/claw-device-pairing -n <namespace> | grep "ClawDevicePairingRequest"

# Check for errors
kubectl logs deployment/claw-device-pairing -n <namespace> | grep -i error
```

### Verify RBAC
```bash
# Check if service account can create CRs
kubectl auth can-i create clawdevicepairingrequests \
  --as=system:serviceaccount:<namespace>:claw-device-pairing \
  -n <namespace>

# Should output "yes"
```

### Test CR creation manually
```bash
cat <<EOF | kubectl apply -f -
apiVersion: pairing.claw.sandbox.redhat.com/v1alpha1
kind: ClawDevicePairingRequest
metadata:
  name: test-manual-request
  namespace: <namespace>
spec:
  requestId: test-manual-request
  selector:
    matchLabels:
      claw.sandbox.redhat.com/instance: "test-instance"
EOF
```

---

## Getting Help

If you encounter issues not covered in this guide:

1. **Check server logs** for detailed error messages
2. **Verify configuration** against `deployment.md`
3. **Test RBAC permissions** using `kubectl auth can-i`
4. **Inspect CRs** using `kubectl get/describe`
5. **Check operator logs** if CRs exist but aren't processed
