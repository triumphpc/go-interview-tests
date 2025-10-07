Задача: "Умный Дом и Устройства"
Время на выполнение: 30-40 минут (включая обсуждение)

**Формулировка для кандидата:**

"Представь, что мы разрабатываем модуль для системы 'Умный Дом'. Этот модуль должен управлять различными устройствами (свет, кондиционер, музыкальный центр) и реагировать на разные события (например, приход домой, уход на работу, наступление ночи).

Сначала система простая, но мы точно знаем, что в будущем:

Количество типов устройств будет расти (жалюзи, сигнализация, чайник).

Количество и сложность сценариев (триггеров для действий) тоже будут расти.

Задача: спроектируй код так, чтобы его было легко расширять новыми устройствами и новыми сценариями. Не нужно писать реализацию методов, достаточно определить ключевые интерфейсы, абстракции и показать, как они будут взаимодействовать. Нарисуй примерную схему и объясни свой выбор."

---



**Что мы проверяем (План для интервьюера):**
1. Постановка задачи и уточняющие вопросы (5 минут)
   Цель: Понимает ли кандидат требования и умеет ли задавать уточняющие вопросы.

Ожидаемые вопросы от кандидата:

"Какие устройства и сценарии есть сейчас?" (Хороший вопрос, показывает, что он думает о конкретике).

"Как именно вызываются сценарии? Вручную пользователем, по таймеру, по датчику?" (Отличный вопрос, разделяющий логику триггера и действия).

"Нужна ли возможность объединять устройства в группы (вся техника в гостиной)?" (Вопрос отличника, показывает системное мышление).

2. Проектирование и написание каркаса кода (15-20 минут)
   Цель: Увидеть, как кандидат применяет принципы SOLID на практике.

Сильное решение (демонстрирует SRP, OCP, DIP):

Интерфейс IDevice: Кандидат выделяет общий интерфейс для всех устройств.

```java
// Пример на Java (можно на любом ЯП)
public interface IDevice {
String getId();
void turnOn();
void turnOff();
// Возможно, boolean isOn();
}
```

Почему хорошо: Позволяет добавлять новые устройства, не меняя существующий код управления ими (OCP).

Конкретные устройства: Реализуют IDevice.

```java
public class Light implements IDevice { ... }
public class AirConditioner implements IDevice { ... }
public class MusicCenter implements IDevice { ... }
```
Интерфейс IScenario (или ICommand, IAction): Кандидат понимает, что сценарий — это набор действий, не зависящий от конкретных устройств.

```java
public interface IScenario {
void execute();
}
```
Конкретные сценарии: Содержат логику, какие устройства включить/выключить. Они зависят от абстракции IDevice, а не от конкретных классов.

```java
public class WelcomeHomeScenario implements IScenario {
private final List<IDevice> devices;

    public WelcomeHomeScenario(List<IDevice> devicesToTurnOn) {
        this.devices = devicesToTurnOn;
    }
    
    @Override
    public void execute() {
        for (IDevice device : devices) {
            device.turnOn();
        }
    }
}
```
Почему хорошо: Принцип Инверсии Зависимостей (DIP). WelcomeHomeScenario не знает о Light или AirConditioner, он знает только об интерфейсе. Добавить новый сценарий (например, GoodNightScenario) легко, не меняя старые (OCP).

Класс-оркестратор (например, SmartHomeController): Управляет устройствами и сценариями. Он получает команду "выполнить сценарий X" и делегирует его выполнение.

```java
public class SmartHomeController {
private final Map<String, IDevice> devices;
private final Map<String, IScenario> scenarios;

    public void executeScenario(String scenarioName) {
        IScenario scenario = scenarios.get(scenarioName);
        if (scenario != null) {
            scenario.execute();
        }
    }
    
    // ... методы для регистрации устройств и сценариев
}
```
Почему хорошо: Контроллер закрыт для модификации (SRP, OCP). Его задача — найти и запустить сценарий, а не знать их внутреннюю логику.

3. Обсуждение и углубление (10-15 минут)
   Цель: Проверить глубину понимания.

Вопросы для обсуждения:

"Как мы можем добавить устройство с уникальным методом, например, setTemperature для кондиционера?"

Плохой ответ: Добавить метод в IDevice (ломает ISP, лампочка будет вынуждена иметь ненужный метод).

Хороший ответ: Создать новый интерфейс, например, ITemperatureDevice, и проверять/приводить тип в конкретном сценарии, который это требует. Или использовать шаблон "Спецификация"/"Посетитель".

Проверяем: Interface Segregation Principle (ISP).

"Представь, что появилось сложное устройство 'Мультимедийная система', которое включает проектор, экран и звук одной командой. Как его интегрировать?"

Хороший ответ: Использовать шаблон "Фасад". MultimediaFacade будет реализовывать IDevice и внутри себя управлять несколькими "дочерними" устройствами.

"А если сценарий должен выполняться не сразу, а по расписанию или при срабатывании датчика?"

Хороший ответ: Выделить отдельную абстракцию для "Триггера" (например, ITrigger). Система будет следить за триггерами, а при их срабатывании — запускать связанный IScenario. Это еще одно применение SRP и DIP.

"Почему ты выбрал именно такую структуру? Какие альтернативы ты рассматривал?"

Проверяем: Способность к рефлексии и понимание компромиссов.

Итог:
Эта задача — отличный индикатор:

Junior: Нарисует класс Light с методом on(), класс HomeController, который напрямую вызывает light.on(). Не видит проблем в жесткой связности.

Middle: Выделит интерфейсы IDevice и IScenario. Покажет, как они взаимодействуют, объяснит OCP и DIP.

Senior: Сразу заговорит о разделении ответственностей (триггеры, команды, исполнители), возможностях для композиции сценариев, проблеме "взрыва" классов при сложной логике и предложит шаблоны вроде "Фасада" или "Посетителя" для решения будущих проблем.

Задание успешно укладывается в 40 минут и дает богатую почву для дискуссии о качестве кода.


ПЛОХОЕ РЕШЕНИЕ
```go
// Плохо: один класс делает всё
type NotificationService struct {
    emailSender *EmailSender
    smsSender   *SMSSender
    pushSender  *PushSender
}

func (n *NotificationService) SendNotification(user *User, message string, notificationType string) error {
    // Нарушает SRP: один метод делает много разных вещей
    switch notificationType {
    case "email":
        // Нарушает OCP: при добавлении нового типа нужно менять этот метод
        return n.emailSender.SendEmail(user.Email, message)
    case "sms":
        return n.smsSender.SendSMS(user.Phone, message)
    case "push":
        return n.pushSender.SendPush(user.DeviceID, message)
    default:
        return fmt.Errorf("unknown notification type")
    }
}

// Конкретные реализации (зависимости)
type EmailSender struct{}
func (e *EmailSender) SendEmail(email, message string) error {
    fmt.Printf("Sending email to %s: %s\n", email, message)
    return nil
}

type SMSSender struct{}
func (s *SMSSender) SendSMS(phone, message string) error {
    fmt.Printf("Sending SMS to %s: %s\n", phone, message)
    return nil
}

type PushSender struct{}
func (p *PushSender) SendPush(deviceID, message string) error {
    fmt.Printf("Sending push to device %s: %s\n", deviceID, message)
    return nil
}
```

ХОРОШЕЕ РЕШЕНИЕ
```go
// SRP: Определяем абстракцию - интерфейс для отправки уведомлений
type Notifier interface {
    Send(user *User, message string) error
    CanHandle(notificationType string) bool
}

// SRP: Каждая структура отвечает только за свой тип уведомлений
type EmailNotifier struct{}

func (e *EmailNotifier) Send(user *User, message string) error {
    fmt.Printf("Sending email to %s: %s\n", user.Email, message)
    return nil
}

func (e *EmailNotifier) CanHandle(notificationType string) bool {
    return notificationType == "email"
}

type SMSNotifier struct{}

func (s *SMSNotifier) Send(user *User, message string) error {
    fmt.Printf("Sending SMS to %s: %s\n", user.Phone, message)
    return nil
}

func (s *SMSNotifier) CanHandle(notificationType string) bool {
    return notificationType == "sms"
}

type PushNotifier struct{}

func (p *PushNotifier) Send(user *User, message string) error {
    fmt.Printf("Sending push to device %s: %s\n", user.DeviceID, message)
    return nil
}

func (p *PushNotifier) CanHandle(notificationType string) bool {
    return notificationType == "push"
}

// OCP: NotificationService закрыт для модификации, но открыт для расширения
type NotificationService struct {
    notifiers []Notifier
}

// OCP: Мы можем добавлять новые типы уведомлений, не изменяя этот метод
func (n *NotificationService) SendNotification(user *User, message string, notificationType string) error {
    for _, notifier := range n.notifiers {
        if notifier.CanHandle(notificationType) {
            return notifier.Send(user, message)
        }
    }
    return fmt.Errorf("no notifier found for type: %s", notificationType)
}

// OCP: Новый тип уведомлений можно легко добавить
type SlackNotifier struct{}

func (s *SlackNotifier) Send(user *User, message string) error {
    fmt.Printf("Sending Slack message to %s: %s\n", user.SlackID, message)
    return nil
}

func (s *SlackNotifier) CanHandle(notificationType string) bool {
    return notificationType == "slack"
}

// DIP: Высокоуровневый модуль зависит от абстракций, а не от конкретных реализаций
type NotificationManager struct {
    service *NotificationService
}

// DIP: Зависимости инжектируются через конструктор
func NewNotificationManager(notifiers ...Notifier) *NotificationManager {
    return &NotificationManager{
        service: &NotificationService{notifiers: notifiers},
    }
}

func (nm *NotificationManager) NotifyUser(user *User, message string, notificationType string) error {
    // Высокоуровневая логика зависит от абстракции Notifier
    return nm.service.SendNotification(user, message, notificationType)
}

// Вспомогательная структура
type User struct {
    Name     string
    Email    string
    Phone    string
    DeviceID string
    SlackID  string
}

// Пример использования
func main() {
    user := &User{
        Name:     "John Doe",
        Email:    "john@example.com",
        Phone:    "+1234567890",
        DeviceID: "device123",
        SlackID:  "U12345",
    }
    
    // DIP: Конкретные реализации инжектируются как зависимости
    notifiers := []Notifier{
        &EmailNotifier{},
        &SMSNotifier{},
        &PushNotifier{},
        &SlackNotifier{}, // OCP: Новый тип добавлен без изменения существующего кода
    }
    
    manager := NewNotificationManager(notifiers...)
    
    // Использование
    manager.NotifyUser(user, "Welcome to our service!", "email")
    manager.NotifyUser(user, "Your code: 123456", "sms")
    manager.NotifyUser(user, "New message received", "slack")
}
```

ЗАДАНИЕ НА SYS DESIGN
--
---

1. "Умный Дом" - Масштабирование до экосистемы
   Задача: Мы успешно сделали ядро системы. Теперь нужно спроектировать архитектуру для:

1 млн пользователей

500 тыс. устройств на пользователя

Интеграция с голосовыми помощниками (Алиса, Siri)

Mobile app + Web portal + устройства

Real-time уведомления

**Что проверяем:**

Микросервисная декомпозиция

Выбор технологий (message queues, DBs)

Security & Authentication

Real-time communication

text
Ожидаемая дискуссия:
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Mobile    │────│  API Gateway│────│ Device Mgmt │
│    App      │    │             │    │  Service    │
└─────────────┘    └─────────────┘    └─────────────┘
│                 │
┌─────────────┐    ┌─────────────┐
│ Auth Service│    │Scenario Exec│
└─────────────┘    └─────────────┘