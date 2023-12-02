chrome.webRequest.onBeforeRequest.addListener(
    function(details) {
        storeVisitedUrl(details.url);
    },
    { urls: ["<all_urls>"] }
);

function storeVisitedUrl(url) {
    chrome.storage.local.get({ visitedUrls: {} }, function (result) {
        let visitedUrls = result.visitedUrls;

        // Check if the URL has already been stored
        if (!visitedUrls[url]) {
            let timestamp = new Date().toISOString();
            visitedUrls[url] = timestamp;
            chrome.storage.local.set({ visitedUrls: visitedUrls });
        }

        // Check for log rotation
        rotateLogsIfNeeded(visitedUrls);
    });
}

function rotateLogsIfNeeded(visitedUrls) {
    let currentDate = new Date().toISOString().split('T')[0];
    chrome.storage.local.get({ lastRotationDate: '' }, function (result) {
        if (currentDate !== result.lastRotationDate) {
            // Rotate the log
            chrome.storage.local.set({ visitedUrls: {}, lastRotationDate: currentDate });

            // Optional: Save the old log somewhere, like downloading it or sending it to a server
            // saveLog(visitedUrls);
        }
    });
}

// Optional: Function to handle saving the old log
// function saveLog(logData) {
//     // Implement logic to save the log
// }
