# Hassukoira - Thefunnydog

## Henkilötiedot

Hassukoira on Telegram-botti, joka tervehtii kaikkia hänelle viestittäviä.

## Ajaminen

Hassukoira on rakennettu Dockerin avulla, joten sen ajaminen on hyvin helppoa. Voit ajaa paikallisen kontin komennolla:

```bash
make compup
```

Komento rakentaa ja ajaa kontin. Mikäli haluat pysäyttää kontin ja siivota jälkesi, aja komento:

```bash
make stop
```

Kontin lokeja ja koiran moikkaajia pääsee lukemaan komennolla:

```bash
make log
```

## To Do

> - [ ] Komennot
>
> `/moi`: Koira vastaa moi-viestillä
>
> `/koira <viesti>`: Koira vastaa viestin sisältöön
>
> `/moikkaajat`: Koira kertoo kolme kovinta moikkaajaa viimeisen kuukauden ajalta

> - [ ] Kielimalli
>
> Jotta Koira voi vastata järkevästi `/koira` komennon viestiin, palvelimella täytyy pyöriä joku äärimmäisen yksinkertainen (kevyt ajaa) kielimalli, jolta pyydetään vastaus viestin sisältöön. Kielimalleja on suhteellisen helppoa ajaa Docker-kontissa, eikä mallin tarvitse olla kovinkaan fiksu, koska tehtävä on niin yksinkertainen.

> - [ ] Lisääminen ryhmään
>
> Koira pitää pystyä lisäämään ryhmään, jolloin sitä käytetään samojen komentojen avulla
