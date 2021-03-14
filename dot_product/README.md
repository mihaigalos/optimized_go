# Dot Product

Simple `result[i] = a[i] * b[i], for 350000000 elements.`
Comparison between pure go implementation (parallelized on 4 cores) and go+rust optimization.


### Results

Each implementation ran 5 times, see [results.csv](https://github.com/mihaigalos/optimized_go/blob/main/dot_product/results.csv).

* Go+Rust is only 10% faster than pure Go, a mean of 364ms (stddev σ = 2.47ms)
* Pure Go requires 403ms on average (stddev σ = 5.31ms).

### Reproducing

```bash
» go run pure_go.go

» cd optimized_with_rust
» make run
```

### Software
```bash
» go version
go version go1.14.3 linux/amd64

» rustc --version
rustc 1.47.0 (18bf6b4f0 2020-10-07)

» cat /etc/os-release
NAME="Ubuntu"
VERSION="18.04.3 LTS (Bionic Beaver)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 18.04.3 LTS"
VERSION_ID="18.04"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=bionic
UBUNTU_CODENAME=bionic
```
### Hardware
```bash
» lscpu
Architecture:        x86_64
CPU op-mode(s):      32-bit, 64-bit
Byte Order:          Little Endian
CPU(s):              4
On-line CPU(s) list: 0-3
Thread(s) per core:  2
Core(s) per socket:  2
Socket(s):           1
NUMA node(s):        1
Vendor ID:           GenuineIntel
CPU family:          6
Model:               60
Model name:          Intel(R) Core(TM) i5-4570T CPU @ 2.90GHz
Stepping:            3
CPU MHz:             2247.844
CPU max MHz:         3600,0000
CPU min MHz:         800,0000
BogoMIPS:            5786.96
Virtualization:      VT-x
L1d cache:           32K
L1i cache:           32K
L2 cache:            256K
L3 cache:            4096K
NUMA node0 CPU(s):   0-3
Flags:               fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm cpuid_fault epb invpcid_single pti ssbd ibrs ibpb stibp tpr_shadow vnmi flexpriority ept vpid fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid xsaveopt dtherm ida arat pln pts md_clear flush_l1d
```
