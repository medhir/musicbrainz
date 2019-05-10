import React, { Component, Fragment } from 'react'
import Search from '../Search'

class Main extends Component {
    render () {
        return (
            <main>
                <header>
                    <h1>Musicbrainz API Challenge</h1>
                </header>
                <Search />
            </main>
        )
    }
}

export default Main