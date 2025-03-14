Rozdělení metod do souborů
cmd/main.go

    Spuštění programu:
        main()
        tea.NewProgram(...)

internal/tui/model.go

    Definice struktury modelu:
        type model struct
        type tickMsg struct
        type frameMsg struct

internal/tui/view.go

    Vykreslování TUI pohledů:
        func (m model) View() string
        func choicesView(m model) string
        func chosenView(m model) string
        func checkbox(label string, checked bool) string
        func progressbar(percent float64) string

internal/tui/update.go

    Obsluha změn v modelu:
        func (m model) Init() tea.Cmd
        func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd)
        func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd)
        func updateChosen(msg tea.Msg, m model) (tea.Model, tea.Cmd)

internal/tui/styles.go

    Stylování a vzhled aplikace:
        var keywordStyle
        var subtleStyle
        var ticksStyle
        var checkboxStyle
        var progressEmpty
        var dotStyle
        var mainStyle

internal/tui/progress.go

    Animace a indikátory pokroku:
        func tick() tea.Cmd
        func frame() tea.Cmd

internal/core/tasks.go

    Business logika kolem úkolů:
        func getTaskName(choice int) string

internal/utils/color.go

    Práce s barvami:
        func makeRampStyles(colorA, colorB string, steps float64) (s []lipgloss.Style)
        func colorToHex(c colorful.Color) string
        func colorFloatToHex(f float64) string
