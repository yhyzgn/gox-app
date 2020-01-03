class X {
}

(function () {
    X.prototype.plugin = function (name, plugin) {
        this[name] = plugin
    };

    window.$ = new X();
})();