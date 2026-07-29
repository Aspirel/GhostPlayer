using desktop.Commands;
using desktop.Views;
using System.ComponentModel;
using System.Windows.Input;

namespace desktop.ViewModels {
    public class MainWindowViewModel: INotifyPropertyChanged {
        public event PropertyChangedEventHandler? PropertyChanged;
        private object _currentView = new HomeViewModel();
        public object CurrentView {
            get => _currentView;
            set {
                if (_currentView != value) {
                    _currentView = value;
                    PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(nameof(CurrentView)));
                }
            }
        }

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

        public ICommand ShowHomeCommand {
            get;
        }

        public ICommand ShowSettingsCommand {
            get;
        }

        public MainWindowViewModel() {
            ShowHomeCommand = new RelayCommand(_ => CurrentView = new HomeViewModel());
            ShowYoutubeCommand = new RelayCommand(_ => CurrentView = new YoutubeViewModel());
            ShowTwitchCommand = new RelayCommand(_ => CurrentView = new TwitchViewModel());
            ShowLocalFilesCommand = new RelayCommand(_ => CurrentView = new LocalFilesViewModel());
            ShowHistoryCommand = new RelayCommand(_ => CurrentView = new HistoryViewModel());
            ShowSettingsCommand = new RelayCommand(_ => CurrentView = new SettingsViewModel());
        }
    }
}
