<?php

namespace CAH\Models;

use MongoDB\Laravel\Relations;

class Card extends BaseModel
{
    public function pack () : Relations\BelongsTo
    {
        return $this->belongsTo (Pack::class);
    }
}
