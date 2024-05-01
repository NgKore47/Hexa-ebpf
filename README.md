# HEXA-UPF
Deployment 

```bash
make clean
```
copy cord directory to home directory
```bash
cp -r cord ~/
```

```bash
DATA_IFACE=eth0 make hexa
```
this will deploy core 
```
make hexaupf
```
this will deploy eupf

to clean only core 
```bash
make hexa-clean
```
to clean everything 
```bash
make clean
```

## 
