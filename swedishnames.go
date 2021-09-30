package namegen

var (
	swedishMaleFirstNames = []string{
		"Abraham", "Adam", "Acke", "Adolf", "Åke", "Albert", "Albin", "Albrecht", "Alexander", "Alf", "Alfred", "Algot", "Alvar", "Anders", "Andreas", "Arne", "Aron", "Arthur", "Arvid", "Asbjörn", "Axel",
		"Bengt", "Bernhard", "Bernt", "Bertil", "Birger", "Bjarne", "Björn", "Björne", "Bo", "Börje", "Bosse", "Bror",
		"Cai", "Caj", "Carl", "Christer", "Christoffer", "Claes",
		"Dag", "Daniel", "Danne",
		"Ebbe", "Eilert", "Einar", "Elias", "Elis", "Elmar", "Elof", "Elov", "Emil", "Emrik", "Enok", "Eric", "Erik", "Erland", "Erling", "Esbjörn", "Eskil", "Evert",
		"Folke", "Frans", "Fredrik", "Frej", "Fritiof", "Fritjof",
		"Gerhard", "Göran", "Gösta", "Göstav", "Göte", "Gottfrid", "Greger", "Gunnar", "Gunne", "Gustaf", "Gustav",
		"Håkan", "Halsten", "Halvar", "Hampus", "Hans", "Harald", "Hasse", "Henrik", "Hilding", "Hjalmar", "Holger",
		"Inge", "Ingemar", "Ingmar", "Ingvar", "Isac", "Isak", "Ivar",
		"Jakob", "Jan", "Janne", "Jarl", "Jens", "Jerk", "Jerker", "Joakim", "Johan", "John", "Jon", "Jonas", "Jöns", "Jöran", "Jörgen",
		"Kalle", "Kåre", "Karl", "Kasper", "Kennet", "Kettil", "Kjell", "Klas", "Knut", "Krister", "Kristian", "Kristofer",
		"Lage", "Lars", "Lasse", "Leif", "Lelle", "Lennart", "Lias", "Loke", "Lorens", "Loui", "Love", "Ludde", "Ludvig",
		"Magnus", "Måns", "Markus", "Mårten", "Martin", "Matheo", "Mats", "Matteus", "Mattias", "Mattis", "Matts", "Melker", "Micael", "Mikael", "Milian", "Müller",
		"Nicklas", "Niklas", "Nils", "Njord", "Noak",
		"Ola", "Oliver", "Olle", "Olaf", "Olof", "Olov", "Örjan", "Orvar", "Östen", "Osvald", "Otto", "Ove",
		"Pål", "Pär", "Patrik", "Peder", "Pehr", "Pelle", "Per", "Peter", "Petter", "Pontus",
		"Ragnar", "Ragnvald", "Rickard", "Rikard", "Robert", "Roffe",
		"Samuel", "Sigfrid", "Sigge", "Sigurður", "Sigvard", "Sivert", "Sixten", "Sören", "Staffan", "Stefan", "Stellan", "Stig", "Sune", "Svante", "Sven",
		"Tage", "Thor", "Thorbjörn", "Thore", "Thorsten", "Thorvald", "Tomas", "Tor", "Torbjörn", "Tore", "Torgny", "Torkel", "Torsten", "Torvald", "Truls", "Tryggve", "Ture",
		"Ulf", "Ulrik", "Uno", "Urban",
		"Valdemar", "Valter", "Verg", "Verner", "Victor", "Vidar", "Vide", "Viggo", "Viktor", "Vilhelm", "Ville", "Vilmar",
		"Yngve",
	}

	swedishFemaleFirstNames = []string{
		"Agda", "Agneta", "Agnetha", "Aina", "Alfhild", "Alicia", "Alva", "Anette", "Anja", "Anneli", "Annika", "Åsa", "Åse", "Aslög", "Asta", "Astrid",
		"Barbro", "Bengta", "Berit", "Birgit", "Birgitta", "Bodil", "Brita", "Britt", "Britta",
		"Cajsa", "Carin", "Carina", "Carita", "Catharina", "Cathrine", "Catrine", "Charlotta", "Christin", "Cilla",
		"Dagny",
		"Ebba", "Edit", "Eira", "Eleonor", "Elin", "Elina", "Ellinor", "Elna", "Elsa", "Elsie", "Embla", "Emelie", "Erica", "Erika", "Erna", "Evy",
		"Fredrika", "Freja", "Frida",
		"Gabriella", "Gerd", "Gerda", "Gertrud", "Gittan", "Göta", "Greta", "Gry", "Gudrun", "Gull", "Gunborg", "Gunda", "Gunhild", "Gunhilda", "Gunilla", "Gunn", "Gunnel", "Gunvor",
		"Hanna", "Hanne", "Hedda", "Hedvig", "Helga", "Henrika", "Hillevi", "Hilma", "Hjördis", "Hulda",
		"Idun", "Ingeborg", "Ingegärd", "Ingegerd", "Inger", "Ingrid",
		"Jannike", "Jennie", "Joline", "Jonna", "Josefin", "Josefina", "Josefine", "Juni",
		"Kaja", "Kajsa", "Kamilla", "Karin", "Karita", "Karla", "Katja", "Katrin", "Kersti", "Kerstin", "Kia", "Kjerstin", "Klara", "Kristin", "Kristine",
		"Laila", "Linn", "Linnéa", "Linnea", "Lis", "Lisbet", "Lisbeth", "Liselott", "Liselotte", "Liv", "Lo", "Lotta", "Lottie", "Lova", "Lovis", "Lovisa",
		"Maj", "Maja", "Majken", "Malena", "Malin", "Margaretha", "Margit", "Mari", "Mariann", "Marit", "Marita", "Märta", "Mathilda", "Meja", "Merit", "Meta", "Mikaela", "Milla", "Milly", "Mimmi", "Minna", "Moa", "Mona", "My",
		"Nanna", "Nea", "Nellie", "Nelly",
		"Ottilia",
		"Pernilla", "Petronella",
		"Ragna", "Ragnhild", "Rakel", "Rebecka", "Rigmor", "Rika", "Ronja", "Runa", "Rut",
		"Saga", "Sanna", "Sassa", "Signe", "Sigrid", "Siri", "Siv", "sofie", "Solveig", "Solvig", "Stina", "Susann", "Susanne", "Svea", "Sylvi", "Synnöve",
		"Tanja", "Tekla", "Terese", "Teresia", "Tessan", "Thea", "Therese", "Thorborg", "Thyra", "Tilde", "Tindra", "Tora", "Torborg", "Tova", "Tove", "Tuva", "Tyra",
		"Ulla", "Ulrica", "Ulrika",
		"Vanja", "Vendela", "Vilhelmina", "Viveka", "Vivi",
		"Ylva",
	}

	swedishLastNames = []string{
		"Åberg", "Abrahamsson", "Abramsson", "Adamsson", "Adolfsson", "Adolvsson", "Ahlberg", "Ahlgren", "Ahlström", "Åkerman", "Åkesson", "Albertsson", "Albinsson", "Albrechtsson", "Albrecktsson", "Albrektson", "Albrektsson", "Alexanderson", "Alexandersson", "Alfredsson", "Alfson", "Alfsson", "Almstedt", "Alvarsson", "Andersson", "Andréasson", "Andreasson", "Arthursson", "Arvidsson", "Axelsson",
		"Beck", "Bengtsdotter", "Bengtsson", "Berg", "Berge", "Bergfalk", "Berggren", "Berglund", "Bergman", "Bergström", "Bernhardsson", "Berntsson", "Björk", "Björklund", "Björkman", "Björnsson", "Blom", "Blomgren", "Blomqvist", "Borg", "Breiner", "Byquist", "Byqvist", "Byström",
		"Carlson", "Carlsson", "Claesson",
		"Dahl", "Dahlman", "Danielsson",
		"Einarsson", "Ek", "Eklund", "Eld", "Eliasson", "Elmersson", "Engberg", "Engman", "Engström", "Ericson", "Ericsson", "Eriksson",
		"Falk", "Feldt", "Forsberg", "Fransson", "Fredriksson", "Frisk",
		"Gerhardsson", "Göransson", "Grahn", "Gunnarsson", "Gustafsson", "Gustavsson",
		"Håkansson", "Hall", "Hallman", "Hansson", "Haraldsson", "Haroldson", "Hellström", "Henriksson", "Herbertsson", "Hermansson", "Hjort", "Holgersson", "Holm", "Holmberg", "Holmström", "Hult",
		"Ingesson", "Isaksson", "Ivarsson",
		"Jakobsson", "Janson", "Jansson", "Johansson", "Johnsson", "Jonasson", "Jönsson", "Jonsson",
		"Karlsson", "Kjellsson", "Klasson", "Knutson", "Knutsson", "Kron",
		"Lager", "Lång", "Larson", "Larsson", "Leifsson", "Lennartsson", "Leonardsson", "Lind", "Lindbeck", "Lindberg", "Lindgren", "Lindholm", "Lindquist", "Lindqvist", "Lindström", "Ljung", "Ljungborg", "Ljunggren", "Ljungman", "Ljungstrand", "Löfgren", "Lund", "Lundberg", "Lundgren", "Lundin", "Lundquist", "Lundqvist", "Lundström",
		"Magnusson", "Mårdh", "Markusson", "Mårtensson", "Martin", "Martinsson", "Matsson", "Mattsson", "Mikaelsson", "Möller",
		"Niklasson", "Nilsson", "Nordström", "Norling", "Nyberg", "Nykvist", "Nylund", "Nyquist", "Nyqvist", "Nyström",
		"Olander", "Oliversson", "Olofsdotter", "Olofsson", "Olson", "Olsson", "Öman", "Östberg", "Ottosson",
		"Patriksson", "Persson", "Petersson", "Pettersson", "Pilkvist",
		"Ragnvaldsson", "Rapp", "Rask", "Robertsson", "Rosenberg", "Rundström",
		"Samuelsson", "Sandberg", "Sandström", "Sigurdsson", "Simonsson", "Sjöberg", "Sjögren", "Söderberg", "Solberg", "Sörensen", "Sorenson", "Sörensson", "Stefansson", "Stenberg", "Stendahl", "Stigsson", "Strand", "Ström", "Sundberg", "Sundén", "Sundström", "Svenson", "Svensson",
		"Tjäder", "Tomasson",
		"Ulfsson",
		"Vång", "Victorsson", "Vinter",
		"Waltersson", "Wang", "Westerberg", "Winter", "Winther", "Wuopio",
	}
)
