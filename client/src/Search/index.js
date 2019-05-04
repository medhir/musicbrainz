import React, { Component, Fragment } from 'react'
import axios from 'axios'
import ResultsList from '../ResultsList'

class Search extends Component {
    constructor (props) {
        super(props)
        this.state = {
            artist: "Pink Floyd",
            title: "Dark Side", 
            filters: {
                album: false, 
                ep: false, 
                single: false
            }
        }
    }

    doSearch () {
        axios.post('/api/search', {
            artist: this.state.artist, 
            title: this.state.title, 
            filters: this.getFiltersArray()
        }).then(response => {
            this.setState({
                results: response.data.ReleaseList.releases
            })
        })
    }

    getFiltersArray () {
        const keys = Object.keys(this.state.filters)
        const filters = []
        keys.forEach(key => {
            if (this.state.filters[key] === true) filters.push(key.toString())
        })
        return filters
    }

    handleArtistChange (e) {
        this.setState({ artist: e.target.value })
    }

    handleTitleChange (e) {
        this.setState({ title: e.target.value })
    }

    handleAlbumFilterChange (e) {
        this.setState({
            filters: {
                album: e.target.checked, 
                ep: this.state.filters.ep,
                single:  this.state.filters.single
            }
        })
    }

    handleEPFilterChange (e) {
        this.setState({
            filters: {
                album: this.state.filters.album, 
                ep: e.target.checked,
                single:  this.state.filters.single
            }
        })
    }

    handleSingleFilterChange (e) {
        this.setState({
            filters: {
                album: this.state.filters.album, 
                ep: this.state.filters.ep,
                single:  e.target.checked
            }
        })
    }

    render () {
        return (
            <Fragment>
                <section className="search">
                    <label>
                        Artist:
                        <input type="text" value={ this.state.artist } onChange={ this.handleArtistChange.bind(this) }/>
                    </label>
                    <label>
                        Title:
                        <input type="text" value={ this.state.title } onChange={ this.handleTitleChange.bind(this) }/>
                    </label>
                    <label>
                        Album:
                        <input type="checkbox" checked={ this.state.filters.album } onChange={ this.handleAlbumFilterChange.bind(this) }/>
                    </label>
                    <label>
                        EP:
                        <input type="checkbox" checked={ this.state.filters.ep } onChange={ this.handleEPFilterChange.bind(this) }/>
                    </label>
                    <label>
                        Single:
                        <input type="checkbox" checked={ this.state.filters.single } onChange={ this.handleSingleFilterChange.bind(this) }/>
                    </label>
                    <button onClick={ this.doSearch.bind(this) }>Search</button>
                </section>
                { this.state.results && <ResultsList results={ this.state.results } /> }
            </Fragment>
        )
    }
}

export default Search