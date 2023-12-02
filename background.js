chrome.webRequest.onBeforeRequest.addListener(
    function(details) {
        console.log("URL:", details.url);
        storeVisitedUrl(details.url);
    },
    { urls: ["<all_urls>"] }
);

function storeVisitedUrl(url) {
    chrome.storage.local.get({ visitedUrls: [] }, function (result) {
        let visitedUrls = result.visitedUrls;
        visitedUrls.push(url);
        chrome.storage.local.set({ visitedUrls: visitedUrls });
    });
}
