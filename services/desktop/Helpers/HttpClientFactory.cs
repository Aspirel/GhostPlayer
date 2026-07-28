using System.Net.Http;

namespace desktop.Helpers {
    internal static class HttpClientFactory {
        private static readonly HttpClient _httpClient;
        static HttpClientFactory() {
            _httpClient = new HttpClient {
                BaseAddress = new Uri("http://localhost:8080")
            };
        }

        public static HttpClient getClient() {
            return _httpClient;
        }

    }
}
