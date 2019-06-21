## [Kubevirt](https://kubevirt.io)

- Bu kubernetes ve VMleri birleştiriyor.
    - CRD ile vmleri pod gibi kullanıyorsun.
    - **Bu aslında işe yarayabilir. Küçük bir CRD modify ile unikernel olabilirler.**
- Sanırım VMlerin deploy edildiği yerde kubelet olması lazım.
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

