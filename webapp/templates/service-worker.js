// Clean up old caches.
self.addEventListener(
  'activate',
  e => {
    e.waitUntil(
      caches.keys()
        .then(cacheNames => Promise.all(
          cacheNames
            .filter(v => '{{ .Version }}' !== v)
            .map(v => caches.delete(v))
        ))
    );
  }
);

// Locally cache eligible components/files.
self.addEventListener(
  'fetch',
  e => {
    return e.respondWith(
      caches.match(e.request)
        .then(r => {
          if (r) {
            return r;
          }
          return fetch(e.request)
            .then(r => {
              if (r.ok) {
                const components = new RegExp('^/((bower_)?components|static)/');
                if (components.test(e.request.url.toString())) {
                  let clone = r.clone();
                  caches.open('{{ .Version }}').then(cache => {
                    cache.put(e.request, clone);
                  });
                }
              }
              return r;
            });
        }
      )
    );
  }
);
