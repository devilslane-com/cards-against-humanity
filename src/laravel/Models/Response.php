<?php

namespace CAH\Models;

use MongoDB\Laravel\Relations;

class Response extends BaseModel
{
    public function cards () : Relations\EmbedsMany
    {
        return $this->embedsMany (Card::class);
    }

    public function round () : Relations\BelongsTo
    {
        return $this->belongsTo (Round::class);
    }
}
