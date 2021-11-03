export function jwtCheck(jwt) {
    if (jwt != undefined) {
        const item = JSON.parse(jwt)
	    const now = new Date()
        if (now.getTime() > item.expiry) {
            localStorage.removeItem('jwt')
            alert('your session is expired')
            return false
        }
        return item.value 
    }
    return false
}