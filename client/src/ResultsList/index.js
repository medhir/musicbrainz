import React, { Component } from 'react'

class ResultsList extends Component {
    render () {
        return (
            <section>
                <ul>
                {
                    this.props.results.map(result => {
                        return <li key={ result.id }>{ result.title }</li>
                    })
                }
                </ul>
            </section>
        )
    }
}

export default ResultsList