# Multiple Eupf deployment with Hexa-core
STEP:1 Deploy Hexa-core without upf 

STEP:2 Deploy eupf
```
git clone https://github.com/evershalik/Hexa-eupf.git ~/
```
```
git checkout -b mutiple_eupf
```
```
cd Hexa-eupf/eupf/docs/deployments/free5gc-ulcl/
```
Edit psp : 
For global-restricted-psp : 
 ```
kubectl get psp global-restricted-psp -o yaml > glrpsp.yaml
sed -i '11a \ \ allowedUnsafeSysctls:\n  - net.ipv4.ip_forward\n  - net.ipv6.conf.all.forwarding' glrpsp.yaml
kubectl apply -f glrpsp.yaml
```
For global-unrestricted-psp : 
```
kubectl get psp global-unrestricted-psp -o yaml > unglrpsp.yaml
sed -i '11a \ \ allowedUnsafeSysctls:\n  - net.ipv4.ip_forward\n  - net.ipv6.conf.all.forwarding' unglrpsp.yaml
kubectl apply -f unglrpsp.yaml

```
Configure iface as required in eupf.yaml && eupf-c.yaml
```
make eupf
```
For Multiple eUPF instances: 
```
helm install \
	edgecomllc-eupf oci://ghcr.io/edgecomllc/charts/eupf \
	--version 0.5.0 \
	--values eupf-c.yaml \
	-n omec-upf-1 \
	--create-namespace
```
Again back to aether-onramp

```
cd
cd aether-onramp
```

At this point the Multiple eUPF(s) will be running (you can verify this using kubectl), but no traffic will be directed to them until UEs are assigned to their IP address pool. Doing so requires loading the appropriate bindings into ROC, which you can do by editing the roc_models line in amp section of vars/main.yml. Comment out the original models file already loaded into ROC, and uncomment the new patch that is to be applied:

Edit these changes in the roc-5g-models-upf2.json file :

```
    cd deps/amp/roles/roc-load/templates   

```
These changes are on the basis of another eupf that have been deployed in another namespace to give Endpoints : 
```
"upf": [
                    {
                        "address": "edgecomllc-eupf.omec-upf-1",
                        "display-name": "Aether UPF 2",
                        "upf-id": "upf-2",
                        "port": 8805,
                        "config-endpoint": "http://edgecomllc-eupf.omec-upf-1.svc:8080"
```
After all changes : 
```
cd vars/main.yml
amp:
   # roc_models: "deps/amp/roles/roc-load/templates/roc-5g-models.json"
   roc_models: "deps/amp/roles/roc-load/templates/roc-5g-models-upf2.json"
```

Then run the following to load the patch: 
```
 make roc-load
```





