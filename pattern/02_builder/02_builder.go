package builder

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Паттерн "Строитель" (Builder) является одним из порождающих паттернов проектирования, который используется
для пошагового создания сложных объектов. Этот паттерн позволяет отделить конструирование объекта от его представления.

Основные компоненты паттерна "Строитель":

1. Строитель (Builder): определяет интерфейс для создания различных частей объекта.
2. Конкретный строитель (ConcreteBuilder): реализует интерфейс Builder и создает конкретный продукт.
3. Продукт (Product): сложный объект, который создается.
4. Директор (Director): Управляет строителем и определяет последовательность шагов для создания продукта.

Плюсы паттерна "Строитель":
1. Разделение процесса конструирования и представления: паттерн "Строитель" позволяет отделить процесс конструирования объекта
от его представления, что упрощает создание различных представлений одного и того же объекта. Возможно создавать разные
представления продукта, не изменяя код.
2. Пошаговое создание объектов: объекты создаются пошагово, можно откладывать шаги конструирования или выполнять их рекурсивно.
3. Инкапсуляция: сложные конструкторы объектов скрыты от клиента, что уменьшает вероятность ошибок.
4. Упрощение тестирования: объекты строителя можно заменять, что упрощает тестирование отдельных частей кода.

Минусы паттерна "Строитель":
1. Усложнение кода: введение дополнительных классов (директор, строители) увеличивает сложность кода.
2. Трудности в поддержке: множество небольших шагов в построении объекта могут затруднять понимание и сопровождение кода.
3. Риск избыточности: если продукт простой, использование паттерна может быть излишним и усложнить проект.

Применение паттерна "Строитель" в реальной практике:

1. Создание объектов с множеством параметров: в случае, если у объекта много параметров или конфигурационных опций,
использование строителя позволяет создавать объекты пошагово, устанавливая каждый параметр в нужный момент.
2. Создание объектов с различными представлениями: если требуется создавать HTML-элементы с различными атрибутами и содержимым,
можно использовать строитель для создания различных видов HTML-элементов.
3. Генерация сложных продуктов в бизнес-логике: для создания заказов с различными комбинациями товаров и услуг.
Строитель позволяет структурировать процесс создания таких объектов.
4. Обработка конфигурационных файлов: если есть конфигурационные файлы с множеством параметров,
строитель может использоваться для пошагового чтения и создания объектов на основе этих файлов.
5. Генерация отчетов и документации: в системах для генерации отчетов или документации часто требуется
создание сложных структур данных или документов с различными разделами и форматированием. Строитель помогает упростить этот процесс.
6. Инициализация объектов с зависимостями: В приложениях с большим количеством зависимостей между объектами
(например, в инъекции зависимостей), строитель может использоваться для инициализации и настройки объектов с учетом их взаимосвязей.
*/

/*
Пример применения паттерна "Строитель"
Создаем конструктор для объектов типа Computer, которые могут иметь различные конфигурации.
*/

import "fmt"

// Определяем структуру для нашего продукта (компьютера)
type Computer struct {
	CPU     string
	GPU     string
	RAM     int
	Storage int
	OS      string
}

func (c Computer) String() string {
	return fmt.Sprintf("Computer [CPU: %s, GPU: %s, RAM: %dGB, Storage: %dGB, OS: %s]",
		c.CPU, c.GPU, c.RAM, c.Storage, c.OS)
}

// Определяем интерфейс строителя, который будет содержать методы для установки различных частей компьютера
type Builder interface {
	SetCPU() Builder
	SetGPU() Builder
	SetRAM() Builder
	SetStorage() Builder
	SetOS() Builder
	Build() Computer
}

// Реализуем конкретного строителя, который будет устанавливать значения для частей компьютера
type GamingComputerBuilder struct {
	computer Computer
}

func (b *GamingComputerBuilder) SetCPU() Builder {
	b.computer.CPU = "Intel i9"
	return b
}

func (b *GamingComputerBuilder) SetGPU() Builder {
	b.computer.GPU = "NVIDIA RTX 3080"
	return b
}

func (b *GamingComputerBuilder) SetRAM() Builder {
	b.computer.RAM = 32
	return b
}

func (b *GamingComputerBuilder) SetStorage() Builder {
	b.computer.Storage = 1000
	return b
}

func (b *GamingComputerBuilder) SetOS() Builder {
	b.computer.OS = "Windows 10"
	return b
}

func (b *GamingComputerBuilder) Build() Computer {
	return b.computer
}

// Создаем директора, который будет управлять процессом строительства, определяя порядок вызова методов строителя
type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) Construct() {
	d.builder.SetCPU().SetGPU().SetRAM().SetStorage().SetOS()
}

// Используем все компоненты для создания компьютера
func main() {
	builder := &GamingComputerBuilder{}
	director := NewDirector(builder)
	director.Construct()
	computer := builder.Build()
	fmt.Println(computer)
}
