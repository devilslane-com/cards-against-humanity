<?php

namespace CAH\Models;

use MongoDB\Laravel\Relations;

class Player extends BaseModel
{
    public function game () : Relations\BelongsTo 
    {
        return $this->belongsTo (Game::class);
    }
}
