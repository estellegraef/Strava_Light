/*
 * 2848869
 */
package activityprocessing

/**
Structure of GPX files may contain the following:
<time> xsd:dateTime </time>                  <!-- Datum und Zeit (UTC/Zulu) in ISO 8601 Format: yyyy-mm-ddThh:mm:ssZ -->
<name> xsd:string </name>                    <!-- Eigenname des Elements -->
<cmt> xsd:string </cmt>                      <!-- Kommentar -->
<desc> xsd:string </desc>                    <!-- Elementbeschreibung -->
<sat> xsd:nonNegativeInteger </sat>          <!-- Anzahl der zur Positionsberechnung herangezogenen Satelliten -->
<hdop> xsd:decimal </hdop>                   <!-- HDOP: Horizontale Streuung der Positionsangabe -->
<vdop> xsd:decimal </vdop>                   <!-- VDOP: Vertikale Streuung der Positionsangabe -->
<pdop> xsd:decimal </pdop>                   <!-- PDOP: Streuung der Positionsangabe -->
<ageofdgpsdata> xsd:decimal </ageofdgpsdata> <!-- Sekunden zwischen letztem DGPS-Empfang und Positionsberechnung -->
*/

/*
important:
- Strecke in km
- wenn Zeitstempel, dann Durchschnittsgeschwindigkeit (ohne Standzeit) und Maximalgeschwindigkeit
- Standzeit (Zeit in der sich nicht bewegt wurde)
- Eingabefehler anhand der Durchschnittsgeschwindigkeit korrigieren
- Datum
 */
func CalculateRouteInKilometers(data string) float64 {
	//TODO implement
	//Harvesine Formula
	return 0
}
