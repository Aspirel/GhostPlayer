using desktop.Commands;
using desktop.Models;
using desktop.Services;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Windows.Input;

namespace desktop.ViewModels {
    public class YoutubeViewModel: INotifyPropertyChanged {
        public event PropertyChangedEventHandler PropertyChanged;
        private readonly YoutubeService _youtubeService;
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


        public YoutubeViewModel() {
            _youtubeService = new YoutubeService();

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
         * Searches for YouTube videos based on the current search query.
         * Clears the existing results and populates the Results collection with new search results.
         */
        private async Task Search() {
            var results = await _youtubeService.SearchAsync(SearchQuery);
            Results.Clear();

            foreach (var result in results) {
                Results.Add(result);
            }
        }

        /**
         * Raises the PropertyChanged event for the SearchQuery property.
         * This method is called whenever the SearchQuery property is updated.
         */
        private void OnSearchQueryChanged() {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(nameof(SearchQuery)));
        }
    }
}
