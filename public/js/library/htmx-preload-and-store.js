(function () {
    const cache = new Map();
    const config = {
        action: "mouseover",
        cacheLifetime: 60 * 5,
        swapType: "innerHTML",
        target: "body",
    };

    function updateConfig() {
        const action = document.body.getAttribute("hx-preload-action")
        const cacheLifetime = document.body.getAttribute("hx-preload-cache-lifetime")
        const swap = document.body.getAttribute("hx-preload-swap")
        const target = document.body.getAttribute("hx-preload-target")

        if (action) {
            config.action = action;
        }

        if (cacheLifetime) {
            const cacheLifetimeSeconds = parseInt(cacheLifetime, 10);
            if (cacheLifetimeSeconds) {
                config.cacheLifetime = cacheLifetimeSeconds;
            }
        }

        if (swap) {
            config.swapType = swap;
        }

        if (target) {
            config.target = target;
        }
    }

    function saveCache(key, value) {
        const isPromise = value instanceof Promise

        cache.set(key, {
            promise: isPromise ? value : null,
            value: isPromise ? null : value,
            cacheLifeTimeEnd: Date.now() + 1000 * config.cacheLifetime
        })
    }

    function getValidCache(key) {
        const cacheItem = cache.get(key);

        if (!cacheItem) {
            return [false, null, null];
        }

        const currentTime = Date.now()
        if (cacheItem && cacheItem.cacheLifeTimeEnd < currentTime) {
            invalidateCache(key);
            return [false, null, null];
        }

        return [true, cacheItem.value, cacheItem.promise];
    }

    function init(node) {
        // preload=init|mouseover
        const preloadAction = node.getAttribute("preload") || config.action;

        if (preloadAction === "init") {
            load(node);
            return
        }

        node.addEventListener(preloadAction, () => load(node));
    }

    function load(node) {
        const hxGet = node.getAttribute('hx-get') || node.getAttribute('data-hx-get');

        // If cache invalidation is false, this cache is good at this moment
        if (cache.get(hxGet) && invalidateCache(hxGet) === false) {
            return;
        }

        const promise = new Promise((resolve, reject) => {
            htmx.ajax('GET', hxGet, {
                source: node,
                handler: function(elt, info) {
                    resolve(info.xhr.responseText);
                    saveCache(hxGet, info.xhr.responseText)
                },
            });
        });

        saveCache(hxGet, promise);
    }

    function invalidateCache(target) {
        const checkForClear = target ? [target] : cache.keys()
        const currentTime = Date.now()
        let invalidateResult = false;

        for (const key of checkForClear) {
            const cacheItem = cache.get(key)
            if (cacheItem && cacheItem.cacheLifeTimeEnd < currentTime) {
                cache.delete(key);
                invalidateResult = true;
            }
        }

        return invalidateResult;
    }

    function swapContent(value, target) {
        htmx.swap(htmx.find(config.target), value, {
            swapStyle: config.swapType,
        }, {
            afterSwapCallback: () => {
                history.pushState({ htmx: true }, '', target)
            }
        });
    }

    function beforeRequest(e) {
        const getHxGet = e.target.getAttribute("hx-get");
        if (!getHxGet) {
            return;
        }

        const cacheInfo = getValidCache(getHxGet);
        const isExist = cacheInfo[0];
        const cacheValue = cacheInfo[1];
        const promiseRequest = cacheInfo[2];

        if (!isExist) {
            return;
        }

        if (promiseRequest) {
            promiseRequest.then((value) => {
                swapContent(value, getHxGet);
            }).catch(() => {
                cache.delete(getHxGet);
               // Something go wrong, need hit directly
            });
        } else {
            swapContent(cacheValue, getHxGet);
        }

        e.preventDefault();
    }

    document.addEventListener("readystatechange", updateConfig);
    htmx.defineExtension("preload-and-store", {
        onEvent: (name, event) => {
            if (name === "htmx:pushedIntoHistory") {
                invalidateCache();
            }

            if (name === "htmx:load") {
                htmx.on("htmx:beforeRequest", beforeRequest);
            }

            if (name === 'htmx:afterProcessNode') {
                const parent = event.target || event.detail.elt;
                parent.querySelectorAll("[preload]").forEach(function (node) {
                    // Initialize the node with the "preload" attribute
                    init(node)

                    // Initialize all child elements that are anchors or have `hx-get` (use with care)
                    node.querySelectorAll('[hx-get],[data-hx-get]').forEach(init)
                })
            }
        }
    });
})();
