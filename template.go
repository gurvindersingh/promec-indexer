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
                    "index": "not_analyzed"
                }
            },
            "dynamic_templates": [
                {
                    "num_fields": {
                        "match": "*num*",
                        "mapping": {
                            "type": "long",
                            "omit_norms": true,
                            "index": "analyzed"
                        }
                    }
                },
                {
                    "scan_fields": {
                        "match": "*scan*",
                        "mapping": {
                            "type": "integer",
                            "omit_norms": true,
                            "index": "analyzed"
                        }
                    }
                },
                {
                    "mass_fields": {
                        "match": "*mass*",
                        "mapping": {
                            "type": "double",
                            "omit_norms": true,
                            "index": "analyzed"
                        }
                    }
                },
                {
                    "charge_fields": {
                        "match": "*_charge",
                        "mapping": {
                            "type": "integer",
                            "omit_norms": true,
                            "index": "analyzed"
                        }
                    }
                },
                {
                    "sec_fields": {
                        "match": "*_sec",
                        "mapping": {
                            "type": "double",
                            "omit_norms": true,
                            "index": "analyzed"
                        }
                    }
                },
                {
                    "xcorr_fields": {
                        "match": "xcorr_*",
                        "mapping": {
                            "type": "double",
                            "omit_norms": true,
                            "index": "analyzed"
                        }
                    }
                },
                {
                    "deltacn_fields": {
                        "match": "deltacn_*",
                        "mapping": {
                            "type": "double",
                            "omit_norms": true,
                            "index": "analyzed"
                        }
                    }
                },
                {
                    "deltacnstar_fields": {
                        "match": "deltacnstar_*",
                        "mapping": {
                            "type": "double",
                            "omit_norms": true,
                            "index": "analyzed"
                        }
                    }
                },
                {
                    "spscore_fields": {
                        "match": "spscore_*",
                        "mapping": {
                            "type": "double",
                            "omit_norms": true,
                            "index": "analyzed"
                        }
                    }
                },
                {
                    "sprank_fields": {
                        "match": "sprank_*",
                        "mapping": {
                            "type": "integer",
                            "omit_norms": true,
                            "index": "analyzed"
                        }
                    }
                },
                {
                    "deltacnstar_fields": {
                        "match": "deltacnstar_*",
                        "mapping": {
                            "type": "double",
                            "omit_norms": true,
                            "index": "analyzed"
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
                                    "index": "not_analyzed"
                                }
                            },
                            "type": "text",
                            "omit_norms": true,
                            "index": "analyzed"
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
