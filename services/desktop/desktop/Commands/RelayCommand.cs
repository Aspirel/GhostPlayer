using System.Windows.Input;

namespace desktop.Commands {
    public class RelayCommand: ICommand {
        private readonly Action<object?> _execute;
        private readonly Func<object?, bool>? _canExecute;
        public event EventHandler? CanExecuteChanged;

        public RelayCommand(Action<object?> execute, Func<object?, bool>? canExecute = null) {
            _execute = execute ?? throw new ArgumentNullException(nameof(execute));
            _canExecute = canExecute;
        }


        /**
         * Determines whether the command can execute in its current state.
         * This method is called by the UI to determine if the command should be enabled or disabled.
         *
         * @param parameter The parameter passed to the command. It can be null.
         * @return True if the command can execute; otherwise, false.
         */
        public bool CanExecute(object? parameter) {
            return _canExecute?.Invoke(parameter) ?? true;
        }

        /**
         * Executes the command with the provided parameter.
         * This method is called when the command is invoked from the UI.
         *
         * @param parameter The parameter passed to the command. It can be null.
         */
        public void Execute(object? parameter) {
            _execute(parameter);
        }
    }
}
