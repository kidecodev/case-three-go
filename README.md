# Case Gin Gorm

silahkan pelajari dari dokumentasi ini
https://gorm.io/docs/update.html

- buat satu buah tabel districts
- dimana berelasi dengan provinces
- wajib migration
- provinces : put | delete
- districts : post| get| put| delete

------------------------------------------------------------------
# Case Point Two
- Tabel sub districs
- Tabel persons
- - id
- - sub district id
- - nip
- - full name
- - first name
- - last name
- - birth date
- - birth place
- - gender [M | F]
- - zona location [WIB | WITA | WIT]
pastikan dapat melakukan :
provinces : post| get all with pagination | put| delete | show one base id
districts : post| get all with pagination | put| delete | show one base id
sub districts : post| get all with pagination | put| delete | show one base id
persons : post| get all with pagination | put| delete | show one base id
please dont show id refrence, but show data from its table
example:
show district data: 
{
"district":"magelang"
"province":"jawa tengah"
}
dont show 
{
"district":"magelang"
"province_id":1
}