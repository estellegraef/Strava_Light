/*
 * 2848869
 * 8089098
 * 3861852
 */

package gps

/**
Structure of GPX files may contain the following:
<ele> xsd:decimal </ele>                     <!-- Höhe in m -->
<time> xsd:dateTime </time>                  <!-- Datum und Zeit (UTC/Zulu) in ISO 8601 Format: yyyy-mm-ddThh:mm:ssZ -->
<magvar> degreesType </magvar>               <!-- Deklination / magnetische Missweisung vor Ort in Grad -->
<geoidheight> xsd:decimal </geoidheight>     <!-- Höhe bezogen auf Geoid -->
<name> xsd:string </name>                    <!-- Eigenname des Elements -->
<cmt> xsd:string </cmt>                      <!-- Kommentar -->
<desc> xsd:string </desc>                    <!-- Elementbeschreibung -->
<src> xsd:string </src>                      <!-- Datenquelle/Ursprung -->
<link> linkType </link>                      <!-- Link zu weiterführenden Infos -->
<sym> xsd:string </sym>                      <!-- Darstellungssymbol -->
<type> xsd:string </type>                    <!-- Klassifikation -->
<fix> fixType </fix>                         <!-- Art der Positionsfeststellung: none, 2d, 3d, dgps, pps -->
<sat> xsd:nonNegativeInteger </sat>          <!-- Anzahl der zur Positionsberechnung herangezogenen Satelliten -->
<hdop> xsd:decimal </hdop>                   <!-- HDOP: Horizontale Streuung der Positionsangabe -->
<vdop> xsd:decimal </vdop>                   <!-- VDOP: Vertikale Streuung der Positionsangabe -->
<pdop> xsd:decimal </pdop>                   <!-- PDOP: Streuung der Positionsangabe -->
<ageofdgpsdata> xsd:decimal </ageofdgpsdata> <!-- Sekunden zwischen letztem DGPS-Empfang und Positionsberechnung -->
<dgpsid> dgpsStationType:integer </dgpsid>   <!-- ID der verwendeten DGPS Station -->
<extensions> extensionsType </extensions>    <!-- GPX Erweiterung -->
*/

func CalculateRouteInKilometers(data string) float64 {
	//TODO implement
	return 0
}
