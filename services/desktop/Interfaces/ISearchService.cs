using desktop.Models;

public interface ISearchService {
    Task<List<SearchResult>> SearchAsync(string searchQuery);
}