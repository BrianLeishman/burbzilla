# Burbzilla

Writing a new dash board/gauge cluster for my 1999 (and possibly my other 1978) Suburban in Golang to monitor all the sensor in my carbed, stroked 383, sbc.

Finna slap.

---
## Static files

JS and CSS are both generated generater from their source languages; TypeScript and SCSS respectively, which can be found in the `ts` and `scss` folders.

TypeScript can be compiled with (while in the project root folder)

```shell
tsc
```

and the SCSS can be compiled with

```shell
node-sass --source-map true --include-path node_modules -o static/css --output-style compressed scss
```