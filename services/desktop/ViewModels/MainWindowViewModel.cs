using desktop.Commands;
using desktop.Views;
using System.ComponentModel;
using System.Windows.Input;

namespace desktop.ViewModels {
    public class MainWindowViewModel: INotifyPropertyChanged {
        public event PropertyChangedEventHandler? PropertyChanged;

        public ICommand ShowYoutubeCommand {
            get;
        }

        public ICommand ShowTwitchCommand {
            get;
        }

        public ICommand ShowHistoryCommand {
            get;
        }

        public ICommand ShowLocalFilesCommand {
            get;
        }

        public ICommand ShowSettingsCommand {
            get;
        }

        public ICommand ShowHomeCommand {
            get;
        }

        private object _currentView;
        public object CurrentView {
            get => _currentView;
            set {
                if (_currentView != value) {
                    _currentView = value;
                    PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(nameof(CurrentView)));
                }
            }
        }

        public MainWindowViewModel() {
            ShowHomeCommand = new RelayCommand(_ => CurrentView = new HomeViewModel());
            ShowYoutubeCommand = new RelayCommand(_ => CurrentView = new YoutubeViewModel());
            ShowTwitchCommand = new RelayCommand(_ => CurrentView = new TwitchViewModel());
            ShowSettingsCommand = new RelayCommand(_ => CurrentView = new SettingsViewModel());
            ShowHistoryCommand = new RelayCommand(_ => CurrentView = new HistoryViewModel());
            ShowLocalFilesCommand = new RelayCommand(_ => CurrentView = new LocalFilesViewModel());
        }
    }
}
