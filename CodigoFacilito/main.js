const url = 'https://jsonplaceholder.typicode.com/posts';
const data = {
    title: 'A new post',
    body: 'Lorem ipsum dolor sit amet',
    userId: 1
};

fetch(url, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
        'Content-Type': 'application/json'
    }
})
    .then(response => response.json())
    .then(json => console.log(json));

    //

    const uri = 'https://jsonplaceholder.typicode.com/posts/1';
    fetch(uri)
        .then(response => response.json())
        .then(data => {

    let element = document.getElementById('elem');
    element.innerHTML = `<h1>${json.title}</h1>`;
    element.innerHTML += `<p>${json.body}</p>`;
    
    console.log(json);
})
.catch(error => console.log(error));