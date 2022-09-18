# Clean Architecture With Go

## Clean Architecture

Uygulamanın daha iyi test edilebilir olmasını, sürdürülebilirliğini ve güvenilirliğini sağlamak amaçlanmıştır. Klasik bilinen n katmanlı mimari yapısındaki karşılaşılan zorlukları ele alarak bunlara çözüm sağlamayı amaçlamaktadır. Katmanlar arası gevşek bir bağımlılık kurmayı hedeflemektedir.

Katmanlı bir mimariye sahiptir. Her katmanın bağımlılığı bir alt katman olacak şekilde iç kısmına doğru olmalıdır. Bu bağımlılık yapısı temel bir ilkedir.

![clean-architecture](https://user-images.githubusercontent.com/16361055/149662671-e2fdab14-7bc3-4585-a482-6b13ac289c5e.png)

## Clean Architecture Katmanları

Uygulamadaki kabaca 4 katmandan oluşmaktadır. Application ve Domain katmanı Core katman olarak isimlendirilmektedir.

### Domain Katmanı

Uygulamanın merkezidir ve tüm uygulamada kullanılacak domain ve veritabanı entitylerini kapsar. Entity, Value Object, Enumeration ve domanin içi Exception sınıflarını içermektedir.

### Application Katmanı

Uygulamanın use case'lerinin bulunduğu ve bu use case'lerin kullandığı arayüzleri içeren katmandır. Custom Exception, Request Response sınıfları, DTO objeleri, Mapping, Validation kontrolleri vb. bu katmanda kullanılmaktadır.

### Infrastructure Katmanı

Application katmanında oluşturulan arayüzlerin implementasyonlarının bulunduğu ve bir takım ayarlamaların yapıldığı katmandır. Database ayarlamaları, bildirim gönderme ve mail gönderimi gibi işlemler gerçekleştirilebilmektedir.

### Presentation Katmanı

Uygulamanın en üst katmanında bulunan ve kullanıcının uygulamayla iletişime geçtiği katmandır.

