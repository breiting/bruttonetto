# Brutto Netto CLI tool

This tool is a CLI frontend for the official [Brutto-Netto-Rechner](https://www.bruttonetto-rechner.at/) for Austria in order to get the official salary numbers
between a certain range. The reason for this tool is to have multiple values retrieved automatically rather than filling all into a web form.

You just need to have Golang installed version >=1.11 (support for GO modules).

```
go run ./...
```

The sample output for 2019 is the following:

```
Brutto/Monat (14x)  Brutto/Jahr  Netto/Jahr  Gesamt/Jahr
3214.29             45000.06     30563.88    58548.90
3571.43             50000.02     33155.68    65054.20
3928.57             54999.98     35747.36    71559.80
4285.71             59999.94     38339.16    78065.10
4642.86             65000.04     40930.96    84570.70
5000.00             70000.00     43522.76    91076.00
5357.14             74999.96     46331.73    97172.19
5714.29             80000.06     49488.85    102612.45
```
