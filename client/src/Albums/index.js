import React, { Component } from 'react'

class Albums extends Component {
    render () {
        return (
            <section className="albums">
                <ul>
                {
                    this.props.albums.map(album => {
                        return <li key={ album.id } onClick={ this.props.onClick } data-release-id={ album.id }>{ album.title }</li>
                    })
                }
                </ul>
            </section>
        )
    }
}

export default Albums