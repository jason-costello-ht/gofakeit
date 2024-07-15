package gofakeit

import (
	"fmt"
	"reflect"
	"testing"
)

func Example_geoShape() {
	_ = Seed(11)
	fmt.Println(GeoShape("connecticut", 5))
	// Output: [[-73.12457018779895 41.95059463095197] [-72.51113419956086 41.99397670944634] [-72.46514213289802 41.65390318732054] [-73.32020620986813 41.854432222521204] [-73.35532430502789 41.873260505567835]]
}

func ExampleFaker_GeoShape() {
	f := New(11)
	fmt.Printf("%v\n", f.GeoShape("connecticut", 5))
	// Output: [[-73.12457018779895 41.95059463095197] [-72.51113419956086 41.99397670944634] [-72.46514213289802 41.65390318732054] [-73.32020620986813 41.854432222521204] [-73.35532430502789 41.873260505567835]]
}

func BenchmarkGeoShape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeoShape("connecticut", 5)
	}
}

func TestFaker_GeoShape(t *testing.T) {
	tests := []struct {
		description string
		state       string
		wanted      [][]float64
	}{
		{
			description: "should return five valid points for Alabama",
			state:       "Alabama",
			wanted:      [][]float64{{-85.71171279391655, 34.93975724524028}, []float64{-86.14528417996435, 31.01524205977796}, []float64{-87.35885561454292, 34.56621922650606}, []float64{-88.32644200650803, 30.99307915706768}, []float64{-86.22558297103113, 34.75811270047316}},
		},
		{
			description: "should return five valid points for Alaska",
			state:       "Alaska",
			wanted:      [][]float64{{-164.52832023957575, 54.711252628832106}, {2.403372901599031, 57.424805536149876}, {-1.140135206054911, 58.27514913002516}, {-118.01690578728444, 56.31501433021281}, {-67.9454224448234, 60.82509714216254}},
		},
		{
			description: "should return five valid points for Arizona",
			state:       "Arizona",
			wanted:      [][]float64{{-110.3708615644019, 36.924113931861996}, []float64{-111.06886306993083, 32.34770567421084}, []float64{-113.02257753463344, 36.48852828810799}, []float64{-111.19813515489861, 36.71229678876395}, []float64{-111.06134846947398, 34.9581682453667}},
		},
		{
			description: "should return five valid points for Arkansas",
			state:       "Arkansas",
			wanted:      [][]float64{{-90.78677138572134, 36.45050391349675}, []float64{-91.38825491438887, 33.62996088228371}, []float64{-93.07181442096564, 36.182042713617406}, []float64{-94.41412458270152, 33.614032438173595}, []float64{-91.49965156498163, 36.31995625838801}},
		},
		{
			description: "should return five valid points for California",
			state:       "California",
			wanted:      [][]float64{{-117.76061398322096, 34.22632458850296}, {-121.2645169319509, 41.14828064580904}, {-122.30803674927897, 40.3191248524881}, {-115.03793005501359, 34.32621759322267}, {-122.49535617981334, 40.48147082760624}},
		},
		{
			description: "should return five valid points for Colorado",
			state:       "Colorado",
			wanted:      [][]float64{{-103.65377064782797, 40.94714204052855}, []float64{-104.50260741894161, 37.71064794148623}, []float64{-106.8785115749641, 40.63909033572698}, []float64{-108.77283148434715, 37.692370498171655}, []float64{-104.6598146718343, 40.797342264128204}},
		},
		{
			description: "should return five valid points for Connecticut",
			state:       "Connecticut",
			wanted:      [][]float64{{-73.12457018779895, 41.95059463095197}, []float64{-72.51113419956086, 41.99397670944634}, []float64{-72.46514213289802, 41.65390318732054}, []float64{-73.32020620986813, 41.854432222521204}, []float64{-73.35532430502789, 41.873260505567835}},
		},
		{
			description: "should return five valid points for Delaware",
			state:       "Delaware",
			wanted:      [][]float64{{-75.26629068049668, 38.70012840056775}, []float64{-75.538627032143, 39.713435425874756}, []float64{-75.61973326826885, 39.59205508432242}, []float64{-75.05467357279032, 38.71475176463182}, []float64{-75.63429243016648, 39.61582095566797}},
		},
		{
			description: "should return five valid points for Florida",
			state:       "Florida",
			wanted:      [][]float64{{-85.25353949530292, 30.400925569312353}, {-82.65038875244139, 28.61878607207447}, {-80.64535272858484, 25.648431373155415}, {-85.08168232014883, 30.53342259003105}, {-81.64137186087699, 28.48603838075419}},
		},
		{
			description: "should return five valid points for Georgia",
			state:       "Georgia",
			wanted:      [][]float64{{-82.45334358963589, 31.187531449679234}, []float64{-84.09640818623645, 34.57930248300268}, []float64{-82.44702384630016, 33.32577598657535}, []float64{-84.58573979524243, 34.173014639130876}, []float64{-84.67357839381947, 34.25256446164111}},
		},
		{
			description: "should return five valid points for Hawaii",
			state:       "Hawaii",
			wanted:      [][]float64{{-156.83086171148574, 20.695290656739427}, {-155.60352501091577, 19.404443447705486}, {-158.22139303745365, 21.14370991504696}, {-164.574421198331, 23.23757046860161}, {-168.72237332821308, 25.07292884792356}},
		},
		{
			description: "should return five valid points for Idaho",
			state:       "Idaho",
			wanted:      [][]float64{{-113.21724566646967, 43.24378354595076}, []float64{-116.98904301892502, 43.2118276091747}, []float64{-115.94082398311495, 47.750506864431024}, []float64{-111.58655432954512, 43.31767241711309}, []float64{-116.05301481417156, 47.87059095713485}},
		},
		{
			description: "should return five valid points for Illinois",
			state:       "Illinois",
			wanted:      [][]float64{{-88.05198187005931, 42.43074111592487}, []float64{-88.59537474907228, 37.961897075124405}, []float64{-90.11633784060538, 42.00539350732944}, []float64{-88.69601282774435, 42.22390253763795}, []float64{-88.58952464653736, 40.511003153092624}},
		},
		{
			description: "should return five valid points for Indiana",
			state:       "Indiana",
			wanted:      [][]float64{{-85.546149716727, 41.705365128462255}, []float64{-85.94676514891829, 38.486063971577565}, []float64{-87.0680924790966, 41.39894985956329}, []float64{-86.02096039316835, 41.55636112020239}, []float64{-85.94245217124428, 40.322409136211}},
		},
		{
			description: "should return five valid points for Iowa",
			state:       "Iowa",
			wanted:      [][]float64{{-91.63339761069872, 43.45717755335498}, []float64{-92.41937792892519, 40.93507372254}, []float64{-94.6193461303218, 43.21712199910602}, []float64{-92.56494396807545, 43.340443031878074}, []float64{-92.41091615912723, 42.37372545119729}},
		},
		{
			description: "should return five valid points for Kansas",
			state:       "Kansas",
			wanted:      [][]float64{{-96.30292189739403, 39.96082553724481}, []float64{-97.20551539543877, 37.531960757742226}, []float64{-99.73188526557287, 39.72964453917667}, []float64{-101.74617221874372, 37.51824423703568}, []float64{-97.37267856406572, 39.848406546280536}},
		},
		{
			description: "should return five valid points for Kentucky",
			state:       "Kentucky",
			wanted:      [][]float64{{-84.63183987940839, 36.9716551532664}, {-84.80220943943465, 39.011527629397705}, {-84.62193627822546, 38.191703173599}, {-83.62005862913715, 38.138426346434834}, {-85.7236558419877, 37.31450023033563}},
		},
		{
			description: "should return five valid points for Louisiana",
			state:       "Louisiana",
			wanted:      [][]float64{{-90.61137238037003, 29.60081231792801}, []float64{-92.40028297081932, 32.64096629972482}, []float64{-92.93304991096488, 32.27679737612498}, []float64{-89.22131158676794, 29.644685771260864}, []float64{-93.02868546879371, 32.34810045268282}},
		},
		{
			description: "should return five valid points for Maine",
			state:       "Maine",
			wanted:      [][]float64{{-70.91202347760056, 43.70979040982699}, []float64{-68.45164138723378, 47.22642680029045}, []float64{-68.35213436648752, 45.82141075164113}, []float64{-68.96026056093878, 44.31805938844386}, []float64{-69.00514390924964, 46.239179816482505}},
		},
		{
			description: "should return five valid points for Maryland",
			state:       "Maryland",
			wanted:      [][]float64{{-76.02024226329439, 39.697107030002314}, []float64{-76.5646425812996, 38.21538640113306}, []float64{-76.6654672408572, 39.62852621010894}, []float64{-76.55878163279735, 39.06058552341547}, []float64{-75.38059743473801, 38.234734327764116}},
		},
		{
			description: "should return five valid points for Massachusetts",
			state:       "Massachusetts",
			wanted:      [][]float64{{-71.1337755053034, 42.27373162119044}, []float64{-72.74175965685946, 42.58372160271575}, []float64{-72.80780071971151, 42.61282750804231}, []float64{-69.98954195786588, 41.28197029900491}, []float64{-71.7013617111863, 42.43004573987364}},
		},
		{
			description: "should return five valid points for Michigan",
			state:       "Michigan",
			wanted:      [][]float64{{-85.03167844248802, 42.879613624579974}, []float64{-87.83980656461841, 47.705616850650955}, []float64{-85.02087749831183, 45.92202861421387}, []float64{-88.67611321691923, 47.12752757930132}, []float64{-82.84964160575345, 42.949259250620955}},
		},
		{
			description: "should return five valid points for Minnesota",
			state:       "Minnesota",
			wanted:      [][]float64{{-94.82794286213732, 48.849765989388615}, []float64{-92.19258116004822, 47.2617934085778}, []float64{-95.6097896099037, 48.3350788584632}, []float64{-95.75013681282674, 48.43585290457602}, []float64{-91.17107885939132, 47.14350878814867}},
		},
		{
			description: "should return five valid points for Mississippi",
			state:       "Mississippi",
			wanted:      [][]float64{{-88.91518691819022, 34.92775491633264}, []float64{-89.34533651895507, 31.009314188492823}, []float64{-89.42500158457236, 34.74639152529723}, []float64{-89.34070557996596, 33.244460679240696}, []float64{-90.9078994592811, 34.13009488450065}},
		},
		{
			description: "should return five valid points for Missouri",
			state:       "Missouri",
			wanted:      [][]float64{{-91.43975863473554, 36.822517701952656}, []float64{-93.6995932923854, 40.19414566252677}, []float64{-91.58928587792013, 40.37634717267485}, []float64{-91.43106659971325, 38.94806361773332}, []float64{-94.37260915031526, 39.790270683649645}},
		},
		{
			description: "should return five valid points for Montana",
			state:       "Montana",
			wanted:      [][]float64{{-106.79824520107934, 48.93592402699845}, []float64{-108.25067640604301, 45.18931104941972}, []float64{-112.31604851720007, 48.57931886114436}, []float64{-108.51967125584012, 48.762513611553544}, []float64{-108.23503970596443, 47.326444013950265}},
		},
		{
			description: "should return five valid points for Nebraska",
			state:       "Nebraska",
			wanted:      [][]float64{{-98.37462350224811, 40.53744717208272}, []float64{-101.33492717999724, 42.729032324111415}, []float64{-98.57049891571016, 42.84746476357005}, []float64{-98.3632372430116, 41.91906902462748}, []float64{-102.21655421083726, 42.466510356286626}},
		},
		{
			description: "should return five valid points for Nevada",
			state:       "Nevada",
			wanted:      [][]float64{{-115.41007809417752, 41.90358269321515}, []float64{-116.1317348235727, 36.255188399826395}, []float64{-118.15166053949352, 41.36596469107367}, []float64{-116.26538793076658, 41.642149124582666}, []float64{-116.12396555381405, 39.47713038137275}},
		},
		{
			description: "should return five valid points for New Hampshire",
			state:       "New Hampshire",
			wanted:      [][]float64{{-71.27005711419181, 43.164077609822876}, []float64{-71.31445091128968, 45.17144971230662}, []float64{-71.26747649790458, 44.36468715201885}, []float64{-71.55455545006524, 43.50146028998331}, []float64{-71.57411367048933, 43.61091041689662}},
		},
		{
			description: "should return five valid points for New Jersey",
			state:       "New Jersey",
			wanted:      [][]float64{{-74.51121530644012, 41.225372436212005}, []float64{-74.47143306087122, 40.430899167221675}, []float64{-74.25034112240024, 40.37926980347294}, []float64{-74.71455787179522, 39.58082170859863}, []float64{-74.73250193615436, 40.667128749756785}},
		},
		{
			description: "should return five valid points for New Mexico",
			state:       "New Mexico",
			wanted:      [][]float64{{-104.39145253751661, 36.920590644442264}, []float64{-105.12287380584986, 32.34715506238639}, []float64{-107.17013058168006, 36.485287941968835}, []float64{-108.80241834052929, 32.32132751348753}, []float64{-105.25833533651743, 36.70891109039145}},
		},
		{
			description: "should return five valid points for New York",
			state:       "New York",
			wanted:      [][]float64{{-73.61162128838485, 44.95215821715024}, {-74.7561784985667, 44.782656802153035}, {-74.56693011827497, 43.37895894349954}, {-73.5151722133895, 43.287738722295806}, {-75.80886257043743, 43.79633605623306}},
		},
		{
			description: "should return five valid points for North Carolina",
			state:       "North Carolina",
			wanted:      [][]float64{{-78.52860166609184, 34.260524678005794}, []float64{-81.54864384125644, 36.33058261052214}, []float64{-78.72842981689257, 36.44244776462837}, []float64{-78.51698563380094, 35.5655331888322}, []float64{-81.34849707027332, 36.38746186673839}},
		},
		{
			description: "should return five valid points for North Dakota",
			state:       "North Dakota",
			wanted:      [][]float64{{-98.27583967071803, 48.957531490312356}, []float64{-99.18220556216558, 46.483941015127}, []float64{-101.71913440602573, 48.72209347017593}, []float64{-103.74184008017116, 46.469971915134806}, []float64{-99.35006738997399, 48.84304238922039}},
		},
		{
			description: "should return five valid points for Ohio",
			state:       "Ohio",
			wanted:      [][]float64{{-82.02715237668919, 39.106494362517246}, []float64{-84.64404822192897, 39.08861745158507}, []float64{-82.02155209987883, 40.91219004795764}, []float64{-83.91678750684126, 41.62766255498837}, []float64{-83.99462618485903, 41.694840451999525}},
		},
		{
			description: "should return five valid points for Oklahoma",
			state:       "Oklahoma",
			wanted:      [][]float64{{-96.39986944383935, 36.954776670443266}, []float64{-97.43649405251388, 34.2222285486605}, []float64{-100.33801858561556, 36.69469088576145}, []float64{-97.62848020645593, 36.82830182293239}, []float64{-97.42533387635922, 35.78092134750998}},
		},
		{
			description: "should return five valid points for Oregon",
			state:       "Oregon",
			wanted:      [][]float64{{-119.35272860198965, 42.76324360187287}, []float64{-124.36615343349328, 42.74361700020688}, []float64{-119.34199964352712, 44.74567054814973}, []float64{-122.97287408894843, 45.53116944959506}, []float64{-117.18523451969666, 42.80862444775497}},
		},
		{
			description: "should return five valid points for Pennsylvania",
			state:       "Pennsylvania",
			wanted:      [][]float64{{-76.73385629239849, 40.22046532334023}, []float64{-80.2811306999424, 40.20772301553871}, []float64{-76.7262649629076, 41.5075294199868}, []float64{-75.20023475721437, 40.24992822865434}, []float64{-75.95830357156049, 41.45132614742418}},
		},
		{
			description: "should return five valid points for Rhode Island",
			state:       "Rhode Island",
			wanted:      [][]float64{{-71.65277370374487, 41.93500227180751}, {-71.39397300686099, 41.971397776625345}, {-71.37456954972497, 41.686092232434916}, {-71.73531000715366, 41.85432657528018}, {-71.75012587657581, 41.870122611111526}},
		},
		{
			description: "should return five valid points for South Carolina",
			state:       "South Carolina",
			wanted:      [][]float64{{-80.20142071410804, 32.60321075717853}, []float64{-81.84468334224327, 34.92648118342113}, []float64{-80.19510020908072, 34.067849804991226}, []float64{-82.33407392829112, 34.64818511834614}, []float64{-82.42192311367803, 34.70267457186028}},
		},
		{
			description: "should return five valid points for South Dakota",
			state:       "South Dakota",
			wanted:      [][]float64{{-98.1871020835132, 45.89666103244943}, []float64{-99.108843090404, 43.100210875265695}, []float64{-101.68880706413137, 45.63049300867088}, []float64{-103.74582485608835, 43.084418490770055}, []float64{-99.27955243818796, 45.76722850594343}},
		},
		{
			description: "should return five valid points for Tennessee",
			state:       "Tennessee",
			wanted:      [][]float64{{-84.68455078951986, 35.28650856693434}, []float64{-87.61716356544274, 36.52425581401481}, []float64{-89.95534995813038, 35.278783349190334}, []float64{-84.87859396972654, 36.59114321091529}, []float64{-84.67327103820314, 36.066810734482104}},
		},
		{
			description: "should return five valid points for Texas",
			state:       "Texas",
			wanted:      [][]float64{{-98.1146773734069, 27.746588785833882}, []float64{-102.56192571018556, 35.53181010669765}, []float64{-98.09757185722886, 32.65455774786395}, []float64{-102.26719393402496, 35.74572566493099}, []float64{-96.3671271413757, 32.44023748361248}},
		},
		{
			description: "should return five valid points for Utah",
			state:       "Utah",
			wanted:      [][]float64{{-110.79873309682715, 37.89361988058392}, []float64{-112.49513384825272, 41.54713550187947}, []float64{-113.84768239680092, 37.87081699916188}, []float64{-110.79220820586076, 40.196873767991846}, []float64{-113.00034983971861, 41.109494231361964}},
		},
		{
			description: "should return five valid points for Vermont",
			state:       "Vermont",
			wanted:      [][]float64{{-71.91821295447933, 44.98451794691385}, []float64{-72.82465400657607, 44.8086607247096}, []float64{-72.20100144018438, 44.89900187232646}, []float64{-72.15424339051351, 44.190814682380335}, []float64{-73.02354828337347, 44.608407010808754}},
		},
		{
			description: "should return five valid points for Virginia",
			state:       "Virginia",
			wanted:      [][]float64{{-78.1500177837831, 37.064598531475504}, []float64{-78.34060713453158, 39.315703424669586}, []float64{-78.13893880395106, 38.41098469072759}, []float64{-75.91181140916981, 37.09541941135991}, []float64{-77.01815636699627, 38.352190953508014}},
		},
		{
			description: "should return five valid points for Washington",
			state:       "Washington",
			wanted:      [][]float64{{-118.73815104585579, 48.95385998988666}, []float64{-119.69751185652592, 46.1629259050073}, []float64{-122.38277410571656, 48.68821698961353}, []float64{-119.87518851087314, 48.82468277248134}, []float64{-119.68718349314986, 47.75492306893167}},
		},
		{
			description: "should return five valid points for West Virginia",
			state:       "West Virginia",
			wanted:      [][]float64{{-78.79130271585763, 39.3299840876418}, []float64{-80.15337680652128, 38.26155805516331}, []float64{-80.20602709364823, 39.71517630472625}, []float64{-80.20197690962657, 38.40578539573067}, []float64{-80.5702050318825, 40.09652633106088}},
		},
		{
			description: "should return five valid points for Wisconsin",
			state:       "Wisconsin",
			wanted:      [][]float64{{-88.577758095276, 43.35439020917421}, []float64{-90.82546171347836, 46.872146528710765}, []float64{-88.56911272004776, 45.57205854777386}, []float64{-91.49486474872063, 46.450767368298955}, []float64{-91.6150274837999, 46.53327202091378}},
		},
		{
			description: "should return five valid points for Wyoming",
			state:       "Wyoming",
			wanted:      [][]float64{{-105.66093121172796, 44.94951232463777}, []float64{-106.50775697497092, 41.712977072920275}, []float64{-108.8780322961923, 44.64145670289666}, []float64{-110.76786430828228, 41.694699397204324}, []float64{-106.66459178284987, 44.79971064350306}},
		},
		{
			description: "should return five valid points for American Samoa",
			state:       "American Samoa",
			wanted:      [][]float64{{-169.1674068492687, -13.95641577477272}, {-169.60371325852543, -13.490186131567787}, {-169.63370772896226, -13.33893683330723}, {-170.38773857210268, -14.267364620634341}, {-170.19918949229725, -12.885381529057776}},
		},
		{
			description: "should return five valid points for Guam",
			state:       "Guam",
			wanted:      [][]float64{{144.84270778945256, 13.678979010375798}, []float64{144.85325974517173, 13.51709000552405}, []float64{144.91190279840427, 13.50656954291738}, []float64{144.78877263031907, 13.343870591863622}, []float64{144.78401309576208, 13.565226265768217}},
		},
		{
			description: "should return five valid points for Northern Mariana Islands",
			state:       "Northern Mariana Islands",
			wanted:      [][]float64{{145.23020525156332, 20.01883086666434}, {145.09500365960446, 19.443361297233483}, {145.8625719348817, 18.111071690384993}, {145.4916799747787, 16.065770597835677}, {145.4773433330891, 18.848449888309233}},
		},
		{
			description: "should return five valid points for Puerto Rico",
			state:       "Puerto Rico",
			wanted:      [][]float64{{-66.16088046915006, 17.963513496136095}, {-67.11894281625496, 18.501114285897753}, {-66.15719545996303, 18.302428458391827}, {-67.05544930389594, 18.515886014061774}, {-65.78440900855331, 18.287628783652366}},
		},
		{
			description: "should return five valid points for U.S. Virgin Islands",
			state:       "Virgin Islands",
			wanted:      [][]float64{{-64.73759000316991, 17.77413952640895}, {-64.95472870340588, 18.388540846211114}, {-65.01939617283085, 18.314943959793396}, {-64.56886385465071, 17.783006152176387}, {-65.031004456014, 18.32935398790559}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			f := New(11)
			if got := f.GeoShape(tt.state, 5); !reflect.DeepEqual(got, tt.wanted) {
				t.Errorf("\ngot : %v\nwant: %v", got, tt.wanted)
			}
		})
	}
}
