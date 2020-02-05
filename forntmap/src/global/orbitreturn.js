export default {
    calculate,
    othenOrbit
}

function othenOrbit(orbitlist) {
    let retorbitlist=new Array()
    for (let index = 0; index < orbitlist.length; index++) {
        const element = orbitlist[index];
        retorbitlist.push(element)
    }
    return retorbitlist
}

function calculate(orbitlist) {
    if (orbitlist.length < 2) {
        return []
    }
    orbitlist.reverse()
    // 计算总时间 
    let averageTime = ( getTime(orbitlist[orbitlist.length -1]) - getTime(orbitlist[0]))/100
    let len = orbitlist.length -1
    let orbitRetuenList =new Array()
    for (let i = 0; i < len; i++) {
        orbitRetuenList = orbitRetuenList.concat(calculateTwoOrbit(averageTime, orbitlist[i], orbitlist[i+1]))
    }
    return orbitRetuenList
}

function calculateTwoOrbit(ti, orbit1, orbit2) {
    let intervalTime =parseInt(( getTime(orbit2) - getTime(orbit1))/ti)+1
    
    let info = getAverageOrbit(orbit1,orbit2)
    let Latitude = info[0]/intervalTime
    let Longitude = info[1]/intervalTime

    let start = getOrbit(orbit1)
    let startLatitude = start[0]
    let startLongitude = start[1]
    let end = getOrbit(orbit1)
    let endLatitude = end[0]
    let endLongitude = end[1]
    let orbitlist = [start]
    
    while (intervalTime > 0) {
        startLatitude +=Latitude
        startLongitude +=Longitude
        orbitlist.push([startLatitude, startLongitude])
        intervalTime --
    }
    orbitlist.push(end)
    return orbitlist
}

function getTime(info) {
    return info.DeviceData.Createtime
}

function getOrbit(info) {
    return [info.DeviceData.Latitude, info.DeviceData.Longitude ] 
}
function getAverageOrbit(info1, info2) {
    return [info2.DeviceData.Latitude - info1.DeviceData.Latitude, info2.DeviceData.Longitude - info1.DeviceData.Longitude] 
}