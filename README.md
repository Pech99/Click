About Click
==========

## Generale
Il programma preme in sequenza i tasti forniti come argomenti da linea di comando. Un tasto è racchiuso in parentesi angolari (<code>\<tasto></code>).

## Tasti permanenti
I tasti permanenti sono: <code>\<ALT></code>, <code>\<SHIFT></code>, <code>\<CTRL></code>, <code>\<RSHIFT></code>, <code>\<RCONTROL></code>, <code>\<Super></code>.
Questi tasti vengono tenuti premuti e rimarranno premuti da qando è stato ricevuto il comanddo a quando non  si riceverà il comando di rilasio (<code>\<-tasto></code>).

## Cmnado Sleep
inserendo la direttiva <code>\<#S[sec]></code> si può efettuare una pausa dalla pressione dei tasti per il numero di milliscecondi desiderato. <code>sec</code> dev'essere un parametro intero positivo.

### Esempio
per scrivere <code>HEY</code>, sarà necessaria la seguente invocazione del'comando:
```
.\Click.exe <SHIFT> <H> <E> <Y>
```
mentre per scrivere <code>Ciao</code>, sarà necessaria la seguente:
```
.\Click.exe <SHIFT> <C> <-SHIFT> <I> <A> <O>
```

## Tasti Mappati
Tasti di Scelta Rapida
<code>\<SP1></code>, <code>\<SP2></code>, <code>\<SP3></code>, <code>\<SP4></code>, <code>\<SP5></code>, <code>\<SP6></code>, <code>\<SP7></code>, <code>\<SP8></code>, <code>\<SP9></code>, <code>\<SP10></code>, <code>\<SP11></code>, <code>\<SP12></code>.

Tasti alfanumerici: <code>\<1></code>, <code>\<2></code>, <code>\<3></code>, <code>\<4></code>, <code>\<5></code>, <code>\<6></code>, <code>\<7></code>, <code>\<8></code>, <code>\<9></code>, <code>\<0></code>, <code>\<Q></code>, <code>\<W></code>, <code>\<E></code>, <code>\<R></code>, <code>\<T></code>, <code>\<Y></code>, <code>\<U></code>, <code>\<I></code>, <code>\<O></code>, <code>\<P></code>, <code>\<A></code>, <code>\<S></code>, <code>\<D></code>, <code>\<F></code>, <code>\<G></code>, <code>\<H></code>, <code>\<J></code>, <code>\<K></code>, <code>\<L></code>, <code>\<Z></code>, <code>\<X></code>, <code>\<C></code>, <code>\<V></code>, <code>\<B></code>, <code>\<N></code>, <code>\<M></code>.

Tasti Funzione: <code>\<F1></code>, <code>\<F2></code>, <code>\<F3></code>, <code>\<F4></code>, <code>\<F5></code>, <code>\<F6></code>, <code>\<F7></code>, <code>\<F8></code>, <code>\<F9></code>, <code>\<F10></code>, <code>\<F11></code>, <code>\<F12></code>, <code>\<F13></code>, <code>\<F14></code>, <code>\<F15></code>, <code>\<F16></code>, <code>\<F17></code>, <code>\<F18></code>, <code>\<F19></code>, <code>\<F20></code>, <code>\<F21></code>, <code>\<F22></code>, <code>\<F23></code>, <code>\<F24></code>.

Tasti Caratteri Speciali: <code>\<MINUS></code>, <code>\<EQUAL></code>, <code>\<SEMICOLON></code>, <code>\<COMMA></code>, <code>\<DOT></code>, <code>\<SLASH></code>, <code>\<GRAVE></code>, <code>\<BACKSLASH></code>, <code>\<APOSTROPHE></code>, <code>\<LEFTBRACE></code>, <code>\<RIGHTBRACE></code>, <code>\<KPASTERISK></code>, 

Tasti White Space: <code>\<TAB></code>, <code>\<ENTER></code>, <code>\<SPACE></code>, 

Tasti Comando: <code>\<ESC></code>, 
<code>\<CAPSLOCK></code>, <code>\<NUMLOCK></code>,
<code>\<SCROLLLOCK></code>, <code>\<RESERVED></code>, <code>\<BACKSPACE></code>,  

Tastierino Nuemrico: <code>\<KP0></code>, <code>\<KP1></code>, <code>\<KP2></code>, <code>\<KP3></code>, <code>\<KP4></code>, <code>\<KP5></code>, <code>\<KP6></code>, <code>\<KP7></code>, <code>\<KP8></code>, <code>\<KP9></code>, <code>\<KPMINUS></code>, <code>\<KPPLUS></code>, <code>\<KPDOT></code>, 

Tasti Funzionalità: <code>\<LBUTTON></code>, <code>\<RBUTTON></code>, <code>\<CANCEL></code>, <code>\<MBUTTON></code>, <code>\<XBUTTON1></code>, <code>\<XBUTTON2></code>, <code>\<BACK></code>, <code>\<CLEAR></code>, <code>\<PAUSE></code>, <code>\<CAPITAL></code>, <code>\<KANA></code>, <code>\<HANGUEL></code>, <code>\<HANGUL></code>, <code>\<JUNJA></code>, <code>\<FINAL></code>, <code>\<HANJA></code>, <code>\<KANJI></code>, <code>\<CONVERT></code>, <code>\<NONCONVERT></code>, <code>\<ACCEPT></code>, <code>\<MODECHANGE></code>, <code>\<PAGEUP></code>, <code>\<PAGEDOWN></code>, <code>\<END></code>, <code>\<HOME></code>, <code>\<LEFT></code>, <code>\<UP></code>, <code>\<RIGHT></code>, <code>\<DOWN></code>, <code>\<SELECT></code>, <code>\<PRINT></code>, <code>\<EXECUTE></code>, <code>\<SNAPSHOT></code>, <code>\<INSERT></code>, <code>\<DELETE></code>, <code>\<HELP></code>, 

Tasti Multimediali: <code>\<SCROLL></code>, <code>\<LMENU></code>, <code>\<RMENU></code>, <code>\<BROWSER_BACK></code>, <code>\<BROWSER_FORWARD></code>, <code>\<BROWSER_REFRESH></code>, <code>\<BROWSER_STOP></code>, <code>\<BROWSER_SEARCH></code>, <code>\<BROWSER_FAVORITES></code>, <code>\<BROWSER_HOME></code>, <code>\<VOLUME_MUTE></code>, <code>\<VOLUME_DOWN></code>, <code>\<VOLUME_UP></code>, <code>\<MEDIA_NEXT_TRACK></code>, <code>\<MEDIA_PREV_TRACK></code>, <code>\<MEDIA_STOP></code>, <code>\<MEDIA_PLAY_PAUSE></code>, <code>\<LAUNCH_MAIL></code>, <code>\<LAUNCH_MEDIA_SELECT></code>, <code>\<LAUNCH_APP1></code>, <code>\<LAUNCH_APP2></code>, <code>\<OEM_1></code>, <code>\<OEM_PLUS></code>, <code>\<OEM_COMMA></code>, <code>\<OEM_MINUS></code>, <code>\<OEM_PERIOD></code>, <code>\<OEM_2></code>, <code>\<OEM_3></code>, <code>\<OEM_4></code>, <code>\<OEM_5></code>, <code>\<OEM_6></code>, <code>\<OEM_7></code>, <code>\<OEM_8></code>, <code>\<OEM_102></code>, <code>\<PROCESSKEY></code>, <code>\<PACKET></code>, <code>\<ATTN></code>, <code>\<CRSEL></code>, <code>\<EXSEL></code>, <code>\<EREOF></code>, <code>\<PLAY></code>, <code>\<ZOOM></code>, <code>\<NONAME></code>, <code>\<PA1></code>, <code>\<OEM_CLEAR></code>, 
