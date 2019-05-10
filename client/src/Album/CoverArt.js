import React, { Component } from 'react'
import axios from 'axios'

const placeholderURL = 'https://upload.wikimedia.org/wikipedia/commons/5/5a/No_CD_Cover_Test.svg'

class CoverArt extends Component {
    constructor (props) {
        super(props)
        this.state = {
            coverURL: null
        }
    }

    componentDidMount () {
        axios.get(`http://coverartarchive.org/release/${ this.props.id }`)
             .then(success => {
                 const img = success.data.images[0]
                 if (img.front === true) {
                     this.setState({
                         coverURL: img.image
                     })
                 }
             })
    }

    render () {
        return (
            <div className="cover">
            {
                this.state.coverURL ? 
                <img src={ this.state.coverURL } alt="Album Cover" /> :
                <img src={ placeholderURL } alt="Album Cover Placeholder" />
            }
            </div>
        )
    }
}

export default CoverArt