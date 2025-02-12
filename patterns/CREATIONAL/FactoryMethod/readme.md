Паттерн `Фабричный метод` относится к порождающим паттернам уровня класса и сфокусирован только на отношениях между классами.

Паттерн `Фабричный метод` полезен, когда система должна оставаться легко расширяемой путем добаления объектов нового типа. Этот паттерн явлется
основой для всех порождающих паттернов и может легко трансформироваться под нужды системы. Поэтому, если перед разработчиком стоят нечеткие
требования для продукта или не ясен способ огранизации взаимодействия между продуктами, то для начала можно воспользоваться паттерном 
`Фабричный метод`, пока полностью не сформируются все требования.

Паттерн `Фабричный метод` применяется для создания объектов с определенным интерфейсом, реализации которого предоставляются потомкам. Другими
словами, есть базовый абстрактный класс фабрики, который говорит, что каждая его наследующая фабрика должна реализовать такой-то метод для
создания своих продуктов.

Реализация `фабричного метода` может быть разной, в большинстве случаев это зависит от языка реализации. Это может быть полиморфизм или параметризированный метод.

Пример: К нам приходят файлы трех расширений `.txt, .png, .doc`. В зависимости от расширения файла мы должны сохранять его в одном из каталогов 
`/file/txt/, /file/png/, /file/doc/`. Значит у нас будет файловая фабрика с параметризированным фабричным методом, принимающим путь к файлу 
который нам нужно сохранить в одном из каталогов. Этот фабричный метод возвращает нам объект, используя который мы можем манипулировать с нашим
файлом (сохранить, посмотреть тип и каталог для сохранения). Заметьте, мы никак не указываем какой экземпляр объект-продукта нам нужно получить,
это делает `фабричный метод` путем определения расширения файла и на его основе выбора подходящего класса продукта. Тем самым, если наша система
будет расширяться и доступных фаловых расширений станет, например 25, то нам всего лишь нужно будет изменить `фабричный метод` и реализовать
классы продуктов.

Требуется для реализации:
	- Базовый абстрактный класс Creator, описывающий интерфейс, который должна реализовать конкретная фабрика для производоства продуктов. Это
	базовый класс описывает `фабричный метод`;
	- Базовый класс Product, описывающий интерфейс продукта, который возвращает фабрика. Все продукты возвращаемые фабрикой должны
	придерживаться единого интерфейса;
	- Класс конкретной фабрики по производству продуктов ConcreteCreator. Этот класс должен реализовать `фабричный метод`;
	- Класс реального продукта productA;

Паттерн `Фабричный метод` отличается от паттерна `Абстрактная фабрика`, тем что `Абстрактная фабрика` производит семейство объектов, эти объекты разные, обладают разными интерфейсами, но взаимодействуют между собой. В то время как `Фабричный метод` производит продукты придерживающиеся одного интерфейса и эти продукты не связаны между собой, не вступают во взаимодействие.

В описании паттерна применяются общие понятия, такие как Класс, Объект, Абстрактный класс. Применимо к языку Go, это Пользовательский Тип,
Значение этого Типа и Интерфейс. Также в языке Go за место общепринятого наследования используется агрегирование и встраивание.