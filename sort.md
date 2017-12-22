Oke, biar saya jelaskan kepada bapak apa itu bubble sort, selection sort dan insertion sort. Tapi pertama saya ingin memberitahu bahwa mereka bertiga itu adalah algoritma untuk mengurutkan sesuatu entah itu dari kecil ke besar atau sebaliknya, ketiganya punya tujuan yang sama tapi dikerjakan dengan cara yang berbeda.

Oke saya akan mencoba menganalogikan masalah ini agar mudah di mengerti oleh bapak. Uji kasusnya adalah 5 presiden indonesia secara berurutan mengantri dalam sebuah barisan: Jokowi, Soekarno, SBY, Soeharto dan Habibie. Sisanya gak ikut dulu hehe. Jadi kita hendak mengurutkan pak presiden dan mantan presiden kita dari yang masa jabatanya paling sebentar hingga paling lama yaitu: Jokowi (3 tahun), Soekarno (22 tahun), SBY (10 tahun), Soeharto (31 tahun) dan Habibie (1 tahun).

# Bubble sort #

Maka bubble sort akan berjalan melewati antrian paling depan yaitu Jokowi dan membandingkannya dengan antrian kedua yaitu dengan Soekarno karna masa jabatan Jokowi lebih kecil dari Soekarno maka bubble sort akan membiarkanya. 

Selanjutnya bubble sort melewati antrian kedua yaitu Soekarno dan membandingkanya dengan antrian ketiga yaitu SBY, karna masa jabatan SBY lebih kecil dari Soekarno maka dia menukar Soekarno dengan SBY dan urutanya pun menjadi Jokowi, SBY, Soekarno, Soeharto dan Habibie.

Sekarang bubble sort melewati antrian ketiga yaitu Soekarno karna pada proses sebelumnya posisi antrian Soekarno ditukar dengan SBY, bubble sort pun membandingkan Soekarno dengan presiden diurutan keempat yaitu Soeharto. Karna pak harto masa jabatanya lebih lama dari Soekarno maka bubble sort akan membiarkanya.

Setelah itu bubble sort melewati antrian keempat yaitu Soeharto dan membandingkanya dengan mantan presiden yang ada di urutan terakhir yaitu Habibie. Karna pak Habibie masa jabatanya lebih kecil dari pak Harto, maka dia akan menukar urutan pak Harto dan pak Habibie dan urutanya pun menjadi Jokowi, SBY, Soekarno, Habibie dan Soeharto. Lalu bubble sort akan kembali ke urutan pertama dan melakukan hal yang sama lagi dari awal. Hingga nanti urutanya menjadi Habibie, Jokowi, SBY, Soekarno dan Soeharto.

# Selection sort #

Berbeda dengan bubble sort, selection sort akan melewati antrian paling depan yaitu Jokowi dan mencari keseluruh antrian siapa yang masa jabatanya paling kecil dan menukar Jokowi dengan orang yang ditemukan. Karena pak Habibie masa jabatanya paling kecil maka urutan antri pak Jokowi ditukar dengan pak Habibie sehingga urutanya menjadi Habibie, Soekarno, SBY, Soeharto dan Jokowi.

Lalu selection sort akan melewati urutan kedua yaitu Soekarno dan mencari presiden atau mantan presiden yang masa jabatanya paling kecil setelah pak Habibie. Karna Jokowi lebih kecil setelah pak Habibie maka urutan antrian Soekarno pun ditukar dengan Jokowi, sehingga urutanya pun menjadi Habibie, Jokowi, SBY, Soeharto dan Soekarno.

Setelah itu dia akan melewati antrian ketiga yaitu SBY dan mencari presiden dengan masa jabatanya paling kecil setelah Habiebie dan Jokowi, karna tidak ada yang masa jabatanya lebih kecil dari SBY maka urutan antrian SBY tidak ditukar dengan siapapun.

Begitu seterusnya hingga urutan antrianya menjadi Habibie, Jokowi, SBY, Soekarno dan Soeharto.

# Insertion sort #

Oke pak, sekarang yang terakhir yaitu insertion sort. Berbeda dengan dua algoritma sebelumnya insertion sort akan berjalan melewati antrian pertama yaitu Jokowi. Setelah itu dia akan berjalan melewati antrian kedua yaitu Soekarno dan membandingkan dengan seluruh antrian yang sudah dilewati insertion sort, karena Soekarno masa jabatanya paling tinggi diantara semua presiden yang dia lewati maka soekarno pun dibiarkan. 

Setelah itu insertion sort akan melewati antrian ke tiga yaitu SBY dan membandingkanya dengan antrian yang sudah dia lewati, karena masa jabatan SBY lebih kecil dari soekarno maka SBY akan dibawa keluar barisan dan menggeser Soekarno ke antrian ke tiga dan memasukan SBY ke antrian kedua. Sehingga urutanya berubah menjadi Jokowi, SBY, Soekarno, Soeharto dan Habibie.

Lalu insertion sort akan berjalan melewati antrian ke empat yaitu Soeharto dan membandingkanya dengan antrian yang sudah dia lewati, karena masa jabatan Soeharto paling tinggi maka insertion sort membiarkannya.

Selanjutnya insertion sort berada pada antrian paling akhir yaitu pak Habibie dan membandingnkannya dengan antrian yang sudah dia lewat, karena masa jabatan pak Habibie lebih kecil dari semua presiden atau mantan presiden maka pak Habibie ditarik keluar dari antrian dan semua presiden atau mantan presiden menggeser antrianya menjadi satu langkah mundur. Pak Soeharto dari antrian ke empat menggeser dirinya menjadi antrian ke lima, Soekarno dari antrian ke empat menggeser dirinya menjadi antrian ke empat, dan seterusnya hingga urutan antrianya menjadi Habibie, Jokowi, SBY, Soekarno dan Soeharto.