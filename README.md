## Interpol - 🚓 for your translations

Maintaing different translations is hard, and we know it.
When the app starts to grow, even using different tools to handle it, there is a chance that a criminal string is going to do something nasty. Intepol is here to help you.

#### Configuration

Create a `.interpol.yml` file inside your project folder.
This file should be like this:

```yaml
master:
  name: "en"
  files:
    - "config/locales/en.yml"
    - "config/locales/something.en.yml"
locales:
  - name: "pt"
    files:
      - "config/locales/pt.yml"
      - "config/locales/something.pt.yml"
```

You can add as many files per language you want.

#### Usage

Run the following command in your project after you configured you manifest, and you interpol will show you if anything is going out of control.

```bash
docker run --rm -v `pwd`:/interpol gregoriokusowski/interpol
```

#### License

MIT
