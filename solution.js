const calculate =()=>{
    const map = new Map();
    for(let i=0;i<list.length; i++) {
        const completer = target - list[i];
        console.log('completer',completer);
        console.log('map.get(completer)',map.get(completer));
        if(map.has(completer)) {
            return [map.get(completer), i];
        }
        map.set(list[i],i)
    }
    return []
}