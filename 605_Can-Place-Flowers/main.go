func canPlaceFlowers(flowerbed []int, n int) bool {
    
    oneFound := false
    for _, v:= range flowerbed {
        if v == 1 {
            oneFound = true
        }
    }
    
    if !oneFound {
        if (len(flowerbed) + 1)/ 2  >= n {
            return true
        }else{
            return false
        }
    }
    
    if len(flowerbed) == 1 && flowerbed[0] == 0  && n == 1 {
        return true
    } 
 //   start1 := true
    nr0 := 0
    canPlace := 0
    last := 0
    lastNr0 := 0
            first1found := false

    for _, v := range flowerbed {
        
        if v == 0  {
            nr0 = nr0 + 1
        }
        
        if v == 1  {
    
            if first1found {
                canPlace = canPlace + (nr0 - 1)/ 2  
            } else {
                
                if nr0 == 0 || nr0 == 1 {
                 canPlace = canPlace   
                }
                if nr0 == 2 {
                    canPlace = canPlace + 1
                    
                }
                if nr0 > 2 {
                     canPlace = canPlace + (nr0 -0) / 2
                }
                first1found = true        
           }
            
        
    //       start1 = false
            nr0 = 0
        }
        
        last  = v
        lastNr0 = nr0
    }
    
    if last == 0 {
                if lastNr0 == 0 || lastNr0 == 1 {
                 canPlace = canPlace   
                }
                if lastNr0 == 2 {
                    canPlace = canPlace + 1
                    
                }
                if lastNr0 > 2 {
                     canPlace = canPlace + (nr0 -0) / 2
                }
    }
    
    if canPlace >= n {
        return true
    }
    return false
}
