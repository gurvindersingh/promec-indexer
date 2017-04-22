package main

const template = `{
    "mappings": {
        "_default_": {
            "_all": {
                "enabled": false
            },
            "properties": {
                "@timestamp": {
                    "type": "date",
                    "format": "strict_date_optional_time||epoch_millis",
                    "index": "true"
                }
            },
            "dynamic_templates": [
                {
                    "num_fields": {
                        "match": "*num*",
                        "mapping": {
                            "type": "long",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "scan_fields": {
                        "match": "*scan*",
                        "mapping": {
                            "type": "integer",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "mass_fields": {
                        "match": "*mass*",
                        "mapping": {
                            "type": "double",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "charge_fields": {
                        "match": "*_charge",
                        "mapping": {
                            "type": "integer",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "sec_fields": {
                        "match": "*_sec",
                        "mapping": {
                            "type": "double",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "xcorr_fields": {
                        "match": "xcorr_*",
                        "mapping": {
                            "type": "double",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "deltacn_fields": {
                        "match": "deltacn_*",
                        "mapping": {
                            "type": "double",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "deltacnstar_fields": {
                        "match": "deltacnstar_*",
                        "mapping": {
                            "type": "double",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "spscore_fields": {
                        "match": "spscore_*",
                        "mapping": {
                            "type": "double",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "index_fields": {
                        "match": "index",
                        "mapping": {
                            "type": "integer",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "sprank_fields": {
                        "match": "sprank_*",
                        "mapping": {
                            "type": "integer",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "deltacnstar_fields": {
                        "match": "deltacnstar_*",
                        "mapping": {
                            "type": "double",
                            "norms": false,
                            "index": true
                        }
                    }
                },
                {
                    "string_fields": {
                        "match": "*",
                        "match_mapping_type": "string",
                        "mapping": {
                            "fields": {
                                "raw": {
                                    "type": "keyword",
                                    "ignore_above": 256,
                                    "index": false
                                }
                            },
                            "type": "text",
                            "norms": false,
                            "index": true
                        }
                    }
                }
            ]
        }
    },
    "settings": {
        "index.refresh_interval": "5s",
        "index.mapper.dynamic": true
    },
    "template": "promec",
    "order"   : 0
}`
