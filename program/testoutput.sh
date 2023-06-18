echo '1-Try passing as arguments --output test00.txt banana standard'
go run main.go --output test00.txt banana standard 

echo '2-Try passing as arguments --output=test00.txt "First\nTest" shadow'
go run main.go --output=test00.txt "First\nTest" shadow | cat -e
cat -e test00.txt

echo '3-Try passing as arguments --output=test01.txt "hello" standard'
go run main.go --output=test01.txt "hello" standard | cat -e
cat -e test01.txt

echo '4-Try passing as arguments --output=test02.txt "123 -> #$%" standard'
go run main.go --output=test02.txt "123 -> #$%" standard | cat -e
cat -e test02.txt

echo '5-Try passing as arguments --output=test03.txt "432 -> #$%&@" shadow'
go run main.go --output=test03.txt "432 -> #$%&@" shadow | cat -e
cat -e test03.txt

echo '6-Try passing as arguments --output=test04.txt "There" shadow'
go run main.go --output=test04.txt "There" shadow | cat -e
cat -e test04.txt

echo '7-Try passing as arguments --output=test05.txt "123 -> \"#$%@" thinkertoy'
go run main.go --output=test05.txt "123 -> \"#$%@" thinkertoy | cat -e
cat -e test05.txt

echo '8-Try passing as arguments --output=test06.txt "2 you" thinkertoy'
go run main.go --output=test06.txt "2 you" thinkertoy | cat -e
cat -e test06.txt

echo '9-Try passing as arguments --output=test07.txt 'Testing long output!' standard'
go run main.go --output=test07.txt 'Testing long output!' standard | cat -e
cat -e test07.txt

echo '10-Try passing as arguments --output=test08.txt 'Line1\n\n\nLine4!' standard'
go run main.go --output=test08.txt 'Line1\n\n\nLine4!' standard | cat -e
cat -e test08.txt

echo '11-Try passing as Zero arguments '
go run main.go 
