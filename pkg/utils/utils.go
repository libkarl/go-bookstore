// budeme dostávat požadavek od uživatele, který bude v JSON formátu 
//(musíme ty data uvolnit, abychom byli schopní je v controlleru použít, co tím kurva myslí ?

package utils 

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r * http.Request, x interface{}) {    // funkce Parse body slouží k analyzování dat od uživatele v JSON formátu, když budu chtít vytvářet novou knihu
	if body, err := ioutil.ReadAll(r.Body); err ==nil{ // do funkce obdržím request a pointer (r *http.Request) r budu pak schopný použít pro přístup k žádosti od uživatele 
		if err := json.Unmarshal([]byte(body), x); err != nil{ //výsledkem je slice bitů, které je potřeba překlopit na objekt formátu Response
			return
		}
	}					
}