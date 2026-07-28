using desktop.Commands;
using desktop.Services;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.DirectoryServices;
using System.Windows.Input;

namespace desktop.ViewModels
{
    public class TwitchViewModel: INotifyPropertyChanged {
        public event PropertyChangedEventHandler PropertyChanged;
        private readonly TwitchService _twitchService;
        private string _searchQuery;
        public string SearchQuery {
            get => _searchQuery;
            set {
                if (_searchQuery != value) {
                    _searchQuery = value;
                    OnSearchQueryChanged();
                }
            }
        }

        public TwitchViewModel() {
            _twitchService = new TwitchService();

            Results = new ObservableCollection<SearchResult>();
            SearchCommand = new RelayCommand(async _ => await Search());
        }

        /**
         * Collection of search results that is bound to the UI.
         * It is updated whenever a new search is performed.
         */
        public ObservableCollection<SearchResult> Results {
            get;
        }

        /**
         * Command that triggers the search operation when executed.
         * It is bound to the UI element responsible for initiating the search.
         */
        public ICommand SearchCommand {
            get;
        }

        /**
         * Searches for Twitch streams based on the current search query.
         * Clears the existing results and populates the Results collection with new search results.
         */
        private async Task Search() {
            var results = await _twitchService.SearchAsync(SearchQuery);
            Results.Clear();

            foreach (var result in results) {
                Results.Add(result);
            }
        }

        /**
         * Invokes the PropertyChanged event for the SearchQuery property.
         * This method is called whenever the SearchQuery property value changes.
         */
        private void OnSearchQueryChanged() {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(nameof(SearchQuery)));
        }
    }
}
