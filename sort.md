Oke, biar saya jelaskan kepada bapak apa itu bubble sort, selection sort dan insertion sort. Tapi pertama saya ingin memberitahu bahwa mereka bertiga itu adalah algoritma untuk mengurutkan sesuatu entah itu dari kecil ke besar atau sebaliknya, ketiganya punya tujuan yang sama tapi dikerjakan dengan cara yang berbeda.

Oke saya akan mencoba menganalogikan masalah ini agar mudah di mengerti oleh bapak. Uji kasusnya adalah 5 presiden indonesia secara berurutan mengantri dalam sebuah barisan: Jokowi, Soekarno, SBY, Soeharto dan Habibie. Sisanya gak ikut dulu hehe. Jadi kita hendak mengurutkan pak presiden dan mantan presiden kita dari yang masa jabatanya paling sebentar hingga paling lama yaitu: Jokowi (3 tahun), Soekarno (22 tahun), SBY (10 tahun), Soeharto (31 tahun) dan Habibie (1 tahun).

# Bubble Short #

Maka bubble short akan berjalan melewati antrian paling depan yaitu Jokowi dan membandingkannya dengan antrian kedua yaitu dengan Soekarno karna masa jabatan Jokowi lebih kecil dari Soekarno maka bubble short akan membiarkanya. 

Selanjutnya bubble short melewati antrian kedua yaitu Soekarno dan membandingkanya dengan antrian ketiga yaitu SBY, karna masa jabatan SBY lebih kecil dari Soekarno maka dia menukar Soekarno dengan SBY dan urutanya pun menjadi Jokowi, SBY, Soekarno, Soeharto dan Habibie.

Sekarang bubble short melewati antrian ketiga yaitu Soekarno karna pada proses sebelumnya posisi antrian Soekarno ditukar dengan SBY, bubble short pun membandingkan Soekarno dengan presiden diurutan keempat yaitu Soeharto. Karna pak harto masa jabatanya lebih lama dari Soekarno maka bubble short akan membiarkanya.

Setelah itu bubble short melewati antrian keempat yaitu Soeharto dan membandingkanya dengan mantan presiden yang ada di urutan terakhir yaitu Habibie. Karna pak Habibie masa jabatanya lebih kecil dari pak Harto, maka dia akan menukar urutan pak Harto dan pak Habibie dan urutanya pun menjadi Jokowi, SBY, Soekarno, Habibie dan Soeharto. Lalu bubble short akan kembali ke urutan pertama dan melakukan hal yang sama lagi dari awal. Hingga nanti urutanya menjadi Habibie, Jokowi, SBY, Soekarno dan Soeharto.

# Selection Short #

Berbeda dengan bubble short, selection short akan melewati antrian paling depan yaitu Jokowi dan mencari keseluruh antrian siapa yang masa jabatanya paling kecil dan menukar Jokowi dengan orang yang ditemukan. Karena pak Habibie masa jabatanya paling kecil maka urutan antri pak Jokowi ditukar dengan pak Habibie sehingga urutanya menjadi Habibie, Soekarno, SBY, Soeharto dan Jokowi.

Lalu selection short akan melewati urutan kedua yaitu Soekarno dan mencari presiden atau mantan presiden yang masa jabatanya paling kecil setelah pak Habibie. Karna Jokowi lebih kecil setelah pak Habibie maka urutan antrian Soekarno pun ditukar dengan Jokowi, sehingga urutanya pun menjadi Habibie, Jokowi, SBY, Soeharto dan Soekarno.

Setelah itu dia akan melewati antrian ketiga yaitu SBY dan mencari presiden dengan masa jabatanya paling kecil setelah Habiebie dan Jokowi, karna tidak ada yang masa jabatanya lebih kecil dari SBY maka urutan antrian SBY tidak ditukar dengan siapapun.

Begitu seterusnya hingga urutan antrianya menjadi Habibie, Jokowi, SBY, Soekarno dan Soeharto.