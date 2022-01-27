async function postData(url, data) {
    const response = await fetch(url, {
        method: 'POST',
        cache: 'no-cache',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });
    if (response.ok) {
        return await response.json();
    } else {
        throw await response.json();
    }
}

async function getData(url) {
    const response = await fetch(url, {
        method: 'GET',
        cache: 'no-cache',
        headers: {
            'Content-Type': 'application/json'
        },
    })
    if (response.ok) {
        return await response.json();
    } else {
        throw await response.json();
    }
}

async function get_db_stats() {
    data = await getData('http://localhost:8080/api/db_stats')
    // console.log(data)
    return data
}
