logrus_honeybadger
==================

This is a simple [logrus](https://github.com/Sirupsen/logrus) hook for sending
your exceptions to the HoneyBadger service.

How to use it?
--------------

Check the logrus documentation for seeing how to configure a hook, but
basically, the needed part for this one is:

    log.AddHook(honeybadger.NewHook("apikey", "live", 5*time.Second))

You can as well check [the
example](https://github.com/agonzalezro/logrus_honeybadger/blob/master/example/example.go)
code in the repo.
