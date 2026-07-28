using desktop.Helpers;
using System.DirectoryServices;
using System.Net.Http;
using System.Net.Http.Json;

namespace desktop.Services {
    public class TwitchService {
        public readonly HttpClient _httpClient;

        public TwitchService() {
            _httpClient = HttpClientFactory.getClient();
        }

        internal async Task<List<SearchResult>> SearchAsync(string searchQuery) {
            try {
                return await _httpClient.GetFromJsonAsync<List<SearchResult>>($"search?q={Uri.EscapeDataString(searchQuery)}");
            } catch (Exception ex) {
                Console.WriteLine($"Error occurred while searching: {ex.Message}");
                return [];
            }
        }
    }
}
