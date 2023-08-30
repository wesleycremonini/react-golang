import { useState } from 'react'

export default function Hello() {
	const [hello, setHello] = useState('Olá')
	const [loading, setLoading] = useState(false)

	async function requestToGo() {
		setLoading(true)
		await fetch('http://localhost:10000/api/hello')
			.then((response) => response.json())
			.then((data) => {
				setHello(data.data.hello)
				setLoading(false)
			})
	}

	return (
		<>
			<button onClick={() => requestToGo()}>Clique</button>
			{loading ? <h1>Carregando...</h1> : <h1>{hello}</h1>}
			<nav>
				<ul>
					<li>
						<a href={`/`}>Voltar</a>
					</li>
				</ul>
			</nav>
		</>
	)
}
