## [Kubevirt](https://kubevirt.io)

- Bu kubernetes ve VMleri birleştiriyor.
    - CRD ile vmleri pod gibi kullanıyorsun.
    - **Bu aslında işe yarayabilir. Küçük bir CRD modify ile unikernel olabilirler.**
- Sanırım VMlerin deploy edildiği yerde kubelet olması lazım.
- **Bu gerçekten promising**
- targetlar:
    - kvm

## Mirage
- **Bu zaten en iyisi**
- Targetlar:
    - xen
    - qubes
    - unix,
    - macosx,
    - virtio,
    - hvt,
    - muen,
    - genode.

## Vorteil.io
- Bu licenced bir ürün o yüzden çok da tartışmamak lazım.
- sanki unik'in etrafını bir wrapperla sarmış gibiler.

## UNIK

- Bu benim yapmak istediğime benzer şeyler yapıyor ama bir kaç sıkıntı var. 
    - Sistem docker olmadan direkt çalışmyor. Docker varsa docker kullanalım zaten. 
    - Image size çok büyük. 

- Sistem olarak o kadar fena değiller , daemon mantığı iyi, o yüzden bundan kesin bahsedeceğim.

## [Virtlet](https://www.mirantis.com/blog/virtlet-run-vms-as-kubernetes-pods/) 
- Bu unikernel support ediyor ama CRI diye bir şey yüklüyorsun. Bir tane script ile çalışıyor.
- Bu [michelangelo](https://github.com/mikelangelo-project/osv-microservice-demo#deploying-unikernels-on-kubernetes) diye bir projeye götürüyor.
- **Buna da docker gerekiyor.**

## Kata containers
- Bu ilginç. 

## ops
- Bu da licensed. Bunlar packaging yapmıyor, kendi runtimelarında çalıştırıyorlar.
- Bu `ops run -n main` yaptığın


## [A vm for every unikernel](http://www.skjegstad.com/blog/2015/03/25/mirageos-vm-per-url-experiment/)

- Bu proje kesinlikle konuşulabilir. Zaten Jitsuya örnek veriyor, bağlaması zor olmasa gerek. 

## [shutit unikernel walkthrough] (https://github.com/ianmiell/shutit-unikernel-walkthrough)

- Bu vagrantta xen ve mirageos çalıştırıyor. Buna bakabilirim lrz cloudda.

## [sysml/minipython] (https://github.com/sysml/minipython)

- Bu da xende python script çalıştırıyor. Buna da bakayım. 

## [Engil/canopy] (https://github.com/engil/canopy)
- Git blogging. Mirage os githubdan kodu alıp run ediyor. *Xen'de* çalışıyor.

- **Bunu kullanmam ama güzel bir örnek.**

## KVM vs Xen
---
`sudo add-apt-repository ppa:avsm/ppa
sudo apt update
sudo apt install opam
opam init
opam install mirage`