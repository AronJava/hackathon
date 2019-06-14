(function (win, lib) {
    var doc = win.document;
    var docEl = doc.documentElement;
    var tid;
    var flexible = lib.flexible || (lib.flexible = {});
    var winW = docEl.clientWidth
    docEl.style.fontSize = winW / 375 * 100 + "px"
    function refreshRem() {
       var winW =docEl.clientWidth
       var rem =(winW / 375)  *100
       docEl.style.fontSize = rem + "px"
       flexible.rem = win.rem = rem;
    }

    win.addEventListener('resize', function () {
        clearTimeout(tid);
        tid = setTimeout(refreshRem, 300);
    }, false);
    win.addEventListener('pageshow', function (e) {
        if (e.persisted) {
            clearTimeout(tid);
            tid = setTimeout(refreshRem, 300);
        }
    }, false);

    refreshRem();

})(window, window['lib'] || (window['lib'] = {}));