using desktop.Helpers;
using desktop.Models;
using System.Net.Http;
using System.Net.Http.Json;

namespace desktop.Services {
    public class YoutubeService: ISearchService {
        public readonly HttpClient _httpClient;

        public YoutubeService() {
            _httpClient = HttpClientFactory.getClient();
        }

        async Task<List<SearchResult>> ISearchService.SearchAsync(string searchQuery)
        {
             if (string.IsNullOrWhiteSpace(searchQuery)) {
                throw new ArgumentException("Search query cannot be empty.");
            }

            try {
                return await _httpClient.GetFromJsonAsync<List<SearchResult>>($"youtube/search?q={Uri.EscapeDataString(searchQuery)}");
            } catch (Exception ex) {
                throw new Exception("An error occurred while searching.", ex);
            }
        }
    }
}
